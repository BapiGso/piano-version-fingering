package core

import (
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func decodeTrackName(trackname string) string {
	detector := chardet.NewTextDetector()
	result, err := detector.DetectBest([]byte(trackname))
	if err != nil {
		// 检测失败，返回原始字符串
		return trackname
	}

	//fmt.Printf("Detected encoding: %s (Confidence: %f)\n", result.Charset, result.Confidence)

	var enc encoding.Encoding
	switch result.Charset {
	case "UTF-8":
		// 已经是 UTF-8，无需转换
		return trackname
	case "Shift_JIS":
		enc = japanese.ShiftJIS
	case "GB-18030": //更广范围的gbk
		enc = simplifiedchinese.GB18030
	case "GBK": //或者 GBK
		enc = simplifiedchinese.GBK
	// 添加其他可能需要处理的编码...
	default:
		// 未知或不支持的编码，返回原始字符串
		return trackname
	}

	decoder := enc.NewDecoder()
	decoded, _, err := transform.String(decoder, trackname)
	if err != nil {
		return trackname // 解码失败
	}
	return decoded
}
