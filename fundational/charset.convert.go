package fundational

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// ToUTF8 convert from CJK encoding to UTF-8
func ToUTF8(from string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(from) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewDecoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewDecoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewDecoder())
	case "shift-jis", "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewDecoder())
	case "cp1252", "windows-1252":
		return fromWindows1252(s), nil
		//return s, nil
	// case "iso-2022-jp", "cp932", "windows-31j":
	// 	reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewDecoder())
	default:
		return s, errors.New("Unsupported encoding " + from)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// FromUTF8 convert from UTF-8 encoding to CJK encoding
func FromUTF8(to string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(to) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewEncoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewEncoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewEncoder())
	case "shift-jis", "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewEncoder())
	// case "iso-2022-jp", "cp932", "windows-31j":
	// 	reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewEncoder())
	default:
		return s, errors.New("Unsupported encoding " + to)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func fromWindows1252(arr []byte) []byte {
	var buf bytes.Buffer
	var r rune

	for _, b := range arr {
		switch b {
		case 0x80:
			r = 0x20AC
		case 0x82:
			r = 0x201A
		case 0x83:
			r = 0x0192
		case 0x84:
			r = 0x201E
		case 0x85:
			r = 0x2026
		case 0x86:
			r = 0x2020
		case 0x87:
			r = 0x2021
		case 0x88:
			r = 0x02C6
		case 0x89:
			r = 0x2030
		case 0x8A:
			r = 0x0160
		case 0x8B:
			r = 0x2039
		case 0x8C:
			r = 0x0152
		case 0x8E:
			r = 0x017D
		case 0x91:
			r = 0x2018
		case 0x92:
			r = 0x2019
		case 0x93:
			r = 0x201C
		case 0x94:
			r = 0x201D
		case 0x95:
			r = 0x2022
		case 0x96:
			r = 0x2013
		case 0x97:
			r = 0x2014
		case 0x98:
			r = 0x02DC
		case 0x99:
			r = 0x2122
		case 0x9A:
			r = 0x0161
		case 0x9B:
			r = 0x203A
		case 0x9C:
			r = 0x0153
		case 0x9E:
			r = 0x017E
		case 0x9F:
			r = 0x0178
		default:
			r = rune(b)
		}

		buf.WriteRune(r)
	}

	return buf.Bytes()
}
