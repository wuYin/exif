package exif

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
)

type TiffReader struct {
	endian binary.ByteOrder
	data   []byte
}

var type2len = []int{1, 1, 2, 4, 8, 1, 1, 2, 4, 8, 4, 8}

func NewTiffReader(endian binary.ByteOrder, data []byte) *TiffReader {
	return &TiffReader{
		endian: endian,
		data:   data,
	}
}

// 解析所有的 IFD tag 数据
func (t *TiffReader) ParseIFDTags(ifdOffset uint32, tag2desc map[uint16]string) (tags []IFDTag, err error) {
	entries := t.read2Bytes(ifdOffset) // 读取 entry 数目
	ifdOffset += 2

	for i := uint16(0); i < entries; i++ {
		start := ifdOffset + 12*uint32(i)
		tagNum := t.read2Bytes(start)      // 标签类别
		tagType := t.read2Bytes(start + 2) // 数据类型
		count := t.read4Bytes(start + 4)   // 组件数量
		offset := start + 8                // tag 数据值或偏移量

		typeLen := uint32(type2len[tagType-1])
		if count*typeLen > 4 {
			offset = t.read4Bytes(offset) // 读取真正的 tag 数据值
		}

		var vs []interface{}
		if tagType == 2 { // ascii string
			v := t.data[offset : offset+count-1] // 直接读取
			vs = append(vs, fmt.Sprintf("%s", v))
		} else {
			for i := uint32(0); i < count; i++ {
				buf := t.readBuf(offset, typeLen)
				var v interface{}
				switch tagType {
				case 1, 6, 7: // unsigned byte / signed byte / undefined
					v = buf
				case 3:
					v = t.endian.Uint16(buf) // unsigned short
				case 4:
					v = t.endian.Uint32(buf) // unsigned long
				case 5, 10:
					up := t.endian.Uint32(buf)
					down := t.endian.Uint32(t.readBuf(offset+4, 4)) // unsigned rational / signed rational
					v = big.NewRat(int64(up), int64(down))
				case 8:
					var x int16
					binary.Read(bytes.NewBuffer(buf), t.endian, &v) // signed short
					v = x
				case 9:
					var x int32
					binary.Read(bytes.NewBuffer(buf), t.endian, &x) // signed long
					v = x
				}
				vs = append(vs, v)
			}
		}

		ifdTag := IFDTag{
			TagNum:  tagNum,
			TagType: tagType,
			Value:   vs,
		}
		if desc, ok := tag2desc[tagNum]; ok {
			ifdTag.Desc = desc
		}
		tags = append(tags, ifdTag)
	}

	return
}

// 解析 exif 数据中的所有 IFD 偏移量
func (t *TiffReader) ParseIFDOffsets() (offsets []uint32, err error) {
	cur := t.read4Bytes(4)
	// 第一个偏移量紧跟头部，一般为 8 字节
	for cur != 0 {
		offsets = append(offsets, cur)
		if cur, err = t.nextOffset(cur); err != nil {
			return nil, err
		}
	}
	return offsets, nil
}

// 获取下一个 IFD 的偏移量
func (t *TiffReader) nextOffset(offset uint32) (next uint32, err error) {
	entries := t.read2Bytes(offset)  // entry 个数
	offset += 2 + 12*uint32(entries) // entry 占 12 字节
	next = t.read4Bytes(offset)
	return
}

// 在指定的偏移量后读取 4 字节内容
func (t *TiffReader) read4Bytes(offset uint32) uint32 {
	return t.endian.Uint32(t.readBuf(offset, 4))
}

// 读取 2 字节内容
func (t *TiffReader) read2Bytes(offset uint32) uint16 {
	return t.endian.Uint16(t.readBuf(offset, 2))
}

// 读取指定偏移量指定长度的字节
func (t *TiffReader) readBuf(offset, size uint32) []byte {
	buf := make([]byte, size)
	copy(buf, t.data[offset:offset+size])
	return buf
}
