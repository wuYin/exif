package exif

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type ExifReader struct {
	Tiff     *TiffReader
	MainTags []IFDTag
	GPSTags  []IFDTag
}

func NewExifReader() *ExifReader {
	return &ExifReader{}
}

type IFDTag struct {
	TagNum  uint16      // 名称
	TagType uint16      // 类型
	Desc    string      // 描述
	Value   interface{} // 指向的值
}

const (
	EXIF_HEADER_FD_SIZE = 12 // 每项 IFD 大小为 12 字节
)

var (
	EXIF_APP_FLAGS = []byte{
		0xFF, 0xD8, // SOI
		0xFF, 0xE1, // APP1
	}
	EXIF_START_FLAG = "Exif"
)

// 读取处理后的 Exif 数据
func (e *ExifReader) ReadContent(img *os.File) error {
	data, endian, err := e.readExifData(img)
	if err != nil {
		log.Fatal("read exif data failed", err)
	}

	t := NewTiffReader(endian, data)
	e.Tiff = t
	offsets, err := t.ParseIFDOffsets()
	if err != nil {
		return fmt.Errorf("parse ifd offsets failed: %v", err)
	}

	for _, offset := range offsets {
		tags, err := t.ParseIFDTags(offset, mainTags)
		if err != nil {
			return fmt.Errorf("parse %v tags failed: %v", offset, err)
		}
		if err = e.SaveTags(tags); err != nil {
			return err
		}
	}
	return nil
}

// 读取 Exif 数据域的数据
func (e *ExifReader) readExifData(r io.Reader) ([]byte, binary.ByteOrder, error) {
	header := make([]byte, EXIF_HEADER_FD_SIZE)
	n, err := r.Read(header)
	if err != nil {
		return nil, nil, err
	}
	if n < EXIF_HEADER_FD_SIZE {
		return nil, nil, errors.New(fmt.Sprintf("read failed: %d bytes", n))
	}
	if !bytes.Equal(EXIF_APP_FLAGS, header[:4]) || string(header[6:10]) != EXIF_START_FLAG {
		return nil, nil, errors.New("img haven't exif header")
	}

	// Exif 数据域最大为 64 KB（2 字节表示）
	var exifLen uint16
	buf := bytes.NewBuffer(header[4:6]) // 数据域大小按大端字节序存储
	if err := binary.Read(buf, binary.BigEndian, &exifLen); err != nil {
		return nil, nil, err
	}

	dataLen := exifLen - 2*8 // 大小描述符本身占用 2 字节
	data := make([]byte, dataLen)
	if _, err := r.Read(data); err != nil {
		return nil, nil, err
	}

	var endian binary.ByteOrder
	switch string(data[:2]) {
	case "II": // Intel 的小端序
		endian = binary.LittleEndian
	case "MM": // Motorola 的大端序
		endian = binary.BigEndian
	default:
		log.Fatalf("unknown byte order: %s", data[:2])
	}

	return data, endian, nil
}

// 存储 tag
func (e *ExifReader) SaveTags(tags []IFDTag) error {
	var err error
	for _, tag := range tags {
		if tag.TagNum == EXIF_IFD_GPS {
			if gpsOffset, ok := tag.Value.([]interface{})[0].(uint32); ok && gpsOffset > 0 {
				if e.GPSTags, err = e.Tiff.ParseIFDTags(gpsOffset, gpsTags); err != nil {
					return err
				}
			}
		}
		e.MainTags = append(e.MainTags, tag)
	}
	return nil
}
