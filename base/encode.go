package base

import (
	"encoding/base64"
	"unsafe"
)

var EncodeTable = map[string]func(Buffer) string {
	"raw": __rawString,
	"HEX": __hexUpperString,
	"hex": __hexLowerString,
	"base64": __base64String,
	"gbk": __gbkString,
}

func __hexUpperString(src Buffer) string {
	return __hexString(src,"0123456789ABCDEF")
}
func __hexLowerString(src Buffer) string {
	return __hexString(src,"0123456789abcdef")
}
func __hexString(src Buffer,tab string) string {
	dst := make([]byte, len(src) * 2)
	for i,v := range src {
		dst[i * 2 + 0] = tab[v >> 4]
		dst[i * 2 + 1] = tab[v & 15]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func __base64String(src Buffer) string {
	return base64.StdEncoding.EncodeToString(src)
}

func __rawString(src Buffer) string {
	return *(*string)(unsafe.Pointer(&src))
}

func __gbkString(src Buffer) string {
	return ""
}
