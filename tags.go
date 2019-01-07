package exif

const (
	EXIF_IFD_GPS = 0x8825 // GPS IFD 的固定偏移量
)

var (
	mainTags = map[uint16]string{
		0x0100: "ImageWidth",                // 缩略图宽
		0x0101: "ImageLength",               // 缩略图长
		0x0102: "BitsPerSample",             //
		0x0103: "Compression",               // 压缩方式	// 1:非压缩 / 6:JPEG 格式压缩
		0x0106: "PhotometricInterpretation", // 图像色彩	// 1:单色 / 2:RGB
		0x0111: "StripOffsets",              //
		0x0115: "SamplesPerPixel",           //
		0x0116: "RowsPerStrip",              //
		0x0117: "StripByteConunts",          //
		0x010e: "ImageDescription",          //
		0x010f: "Make",                      //
		0x0110: "Model",                     //
		0x0112: "Orientation",               //
		0x011a: "XResolution",               //
		0x011b: "YResolution",               //
		0x011c: "PlanarConfiguration",       //
		0x0128: "ResolutionUnit",            // 分辨率单位	// 1:英寸 / 2:厘米
		0x0131: "Software",                  //
		0x0132: "DateTime",                  //
		0x013e: "WhitePoint",                //
		0x013f: "PrimaryChromaticities",     //
		0x0201: "JpegIFOffset",              //
		0x0202: "JpegIFByteCount",           // JPEG 图像文件大小
		0x0211: "YCbCrCoefficients",         //
		0x0213: "YCbCrPositioning",          //
		0x0214: "ReferenceBlackWhite",       //
		0x8298: "Copyright",                 //

		// 子 IFD 数据
		0x8769: "ExifOffset",
		0x829a: "ExposureTime",               //
		0x829d: "FNumber",                    //
		0x8822: "ExposureProgram",            //
		0x8827: "ISOSpeedRatings",            //
		0x9000: "ExifVersion",                //
		0x9003: "DateTimeOriginal",           // 照片拍摄的日期时间，格式为 YYYY:MM:DD HH:MM:SS
		0x9004: "DateTimeDigitized",          //
		0x9101: "ComponentConfiguration",     //
		0x9102: "CompressedBitsPerPixel",     //
		0x9201: "ShutterSpeedValue",          //
		0x9202: "ApertureValue",              //
		0x9203: "BrightnessValue",            //
		0x9204: "ExposureBiasValue",          //
		0x9205: "MaxApertureValue",           //
		0x9206: "SubjectDistance",            //
		0x9207: "MeteringMode",               //
		0x9208: "LightSource",                //
		0x9209: "Flash",                      //
		0x920a: "FocalLength",                //
		0x927c: "MakerNote",                  //
		0x9286: "UserComment",                //
		0xa000: "FlashPixVersion",            //
		0xa001: "ColorSpace",                 //
		0xa002: "ExifImageWidth",             // 原图宽度
		0xa003: "ExifImageHeight",            // 原图高度
		0xa004: "RelatedSoundFile",           //
		0xa005: "ExifInteroperabilityOffset", //
		0xa20e: "FocalPlaneXResolution",      //
		0xa20f: "FocalPlaneYResolution",      //
		0xa210: "FocalPlaneResolutionUnit",   //
		0xa217: "SensingMethod",              //
		0xa300: "FileSource",                 //
		0xa301: "SceneType",                  //
	}

	gpsTags = map[uint16]string{
		0x0000: "GPSVersionID",
		0x0001: "GPSLatitudeRef",
		0x0002: "GPSLatitude",
		0x0003: "GPSLongitudeRef",
		0x0004: "GPSLongitude",
		0x0005: "GPSAltitudeRef",
		0x0006: "GPSAltitude",
		0x0007: "GPSTimeStamp",
		0x0008: "GPSSatellites",
		0x0009: "GPSStatus",
		0x000a: "GPSMeasureMode",
		0x000b: "GPSDOP",
		0x000c: "GPSSpeedRef",
		0x000d: "GPSSpeed",
		0x000e: "GPSTrackRef",
		0x000f: "GPSTrack",
		0x0010: "GPSImgDirectionRef",
		0x0011: "GPSImgDirection",
		0x0012: "GPSMapDatum",
		0x0013: "GPSDestLatitudeRef",
		0x0014: "GPSDestLatitude",
		0x0015: "GPSDestLongitudeRef",
		0x0016: "GPSDestLongitude",
		0x0017: "GPSDestBearingRef",
		0x0018: "GPSDestBearing",
		0x0019: "GPSDestDistanceRef",
		0x001a: "GPSDestDistance",
		0x001b: "GPSProcessingMethod",
		0x001c: "GPSAreaInformation",
		0x001d: "GPSDateStamp",
		0x001e: "GPSDifferential",
	}
)
