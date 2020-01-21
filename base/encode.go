package base

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base32"
	"encoding/base64"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"unsafe"
)

var EncodeTable = map[string]func(Buffer) string{
	"raw":     __rawString,
	"utf8":    __rawString,
	"HEX":     __HEXString,
	"hex":     __hexString,
	"BCD":     __BCDString,
	"bcd":     __bcdString,
	"base32":  __base32String,
	"base64":  __base64String,
	"gbk":     __gbkString,
	"gb18030": __gb18030String,
	"md5":     __md5String,
	"MD5":     __MD5String,
	"sha1":    __sha1String,
	"SHA1":    __SHA1String,
}

var __number = []func(Buffer) uint64{
	func(n Buffer) (r uint64) {
		for i := len(n) - 1; i >= 0; i-- {
			r = (r << 8) + uint64(n[i])
		}
		return
	},
	func(n Buffer) (r uint64) {
		for _, v := range n {
			r = (r << 8) + uint64(v)
		}
		return
	},
	func(n Buffer) (r uint64) {
		for _, v := range n {
			r = r*100 + uint64(v>>4)*10 + uint64(v&0x0f)
		}
		return
	},
	func(n Buffer) (r uint64) {
		for i := len(n) - 1; i >= 0; i-- {
			r = r*100 + uint64(n[i]>>4)*10 + uint64(n[i]&0x0f)
		}
		return
	},
}

func __HEXString(src Buffer) string {
	return __hex(src, "0123456789ABCDEF")
}

func __hexString(src Buffer) string {
	return __hex(src, "0123456789abcdef")
}

func __hex(src Buffer, tab string) string {
	dst := make([]byte, len(src)*2)
	for i, v := range src {
		dst[i*2+0] = tab[v>>4]
		dst[i*2+1] = tab[v&15]
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func __base32String(src Buffer) string {
	return base32.StdEncoding.EncodeToString(src)
}

func __base64String(src Buffer) string {
	return base64.StdEncoding.EncodeToString(src)
}

func __rawString(src Buffer) string {
	return *(*string)(unsafe.Pointer(&src))
}

func __gb18030String(src Buffer) string {
	dst, _, _ := transform.Bytes(simplifiedchinese.GB18030.NewDecoder(), src)
	return *(*string)(unsafe.Pointer(&dst))
}

func __gbkString(src Buffer) string {
	dst, _, _ := transform.Bytes(simplifiedchinese.GBK.NewDecoder(), src)
	return *(*string)(unsafe.Pointer(&dst))
}

func __md5String(src Buffer) string {
	h := md5.New()
	h.Write(src)
	dst := h.Sum(nil)
	return __hexString(dst)
}

func __MD5String(src Buffer) string {
	h := md5.New()
	h.Write(src)
	dst := h.Sum(nil)
	return __HEXString(dst)
}

func __sha1String(src Buffer) string {
	h := sha1.New()
	h.Write(src)
	dst := h.Sum(nil)
	return __hexString(dst)
}

func __SHA1String(src Buffer) string {
	h := sha1.New()
	h.Write(src)
	dst := h.Sum(nil)
	return __HEXString(dst)
}

func __BCDString(src Buffer) string {
	dst := make([]byte, len(src)*2)
	for i, v := range src {
		dst[i*2+0] = 48 + (v >> 4)
		dst[i*2+1] = 48 + (v & 0x0f)
	}
	return *(*string)(unsafe.Pointer(&dst))
}

func __bcdString(src Buffer) string {
	dst := make([]byte, len(src)*2)
	for i := len(src) - 1; i >= 0; i-- {
		dst[i*2+0] = 48 + (src[i] >> 4)
		dst[i*2+1] = 48 + (src[i] & 0x0f)
	}
	return *(*string)(unsafe.Pointer(&dst))
}
