package io

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/vlorc/lua-vm/base"
	"io"
	"io/ioutil"
	"strings"
	"unsafe"
)

type Reader interface {
	io.ReadSeeker
	Endian(bool)
	Reset(io.Reader)
	ReadStringAt(int) (string, error)
	ReadBufferAt(int) (base.Buffer, error)
	ReadLine() (string, error)
	ReadInt8() (uint8, error)
	ReadInt16() (uint16, error)
	ReadInt32() (uint32, error)
	ReadInt64() (uint64, error)
	ReadRune() (rune, error)
}

type StreamReader struct {
	reader *bufio.Reader
	order  binary.ByteOrder
}
type ReaderFactory struct{}

func (f ReaderFactory) New(b interface{}) Reader {
	switch r := b.(type) {
	case *[]byte:
		if nil != r {
			return f.FormBuffer(base.Buffer(*r))
		}
	case *base.Buffer:
		if nil != r {
			return f.FormBuffer(base.Buffer(*r))
		}
	case *string:
		if nil != r {
			return f.FormString(*r)
		}
	case []byte:
		return f.FormBuffer(base.Buffer(r))
	case base.Buffer:
		return f.FormBuffer(base.Buffer(r))
	case string:
		return f.FormString(r)
	case io.Reader:
		return f.FormStream(r)
	}
	return nil
}

func (ReaderFactory) FormStream(r io.Reader) Reader {
	return &StreamReader{
		reader: bufio.NewReader(r),
		order:  binary.LittleEndian,
	}
}
func (f ReaderFactory) FormString(str string) Reader {
	return f.FormStream(strings.NewReader(str))
}
func (f ReaderFactory) FormBuffer(buf base.Buffer) Reader {
	return f.FormStream(bytes.NewReader(buf))
}
func (f ReaderFactory) ReadBuffer(r io.Reader) (base.Buffer, error) {
	buf, err := ioutil.ReadAll(r)
	return base.Buffer(buf), err
}
func (f ReaderFactory) ReadString(r io.Reader) (string, error) {
	buf, err := f.ReadBuffer(r)
	if nil != err {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&buf)), nil
}
func (r *StreamReader) Endian(big bool) {
	if big {
		r.order = binary.BigEndian
	} else {
		r.order = binary.LittleEndian
	}
}
func (r *StreamReader) ReadStringAt(delim int) (str string, err error) {
	return r.reader.ReadString(byte(delim))
}
func (r *StreamReader) ReadLine() (str string, err error) {
	val, _, err := r.reader.ReadLine()
	if len(val) > 0 {
		str = string(val)
	}
	return
}
func (r *StreamReader) ReadBufferAt(delim int) (base.Buffer, error) {
	val, err := r.reader.ReadBytes(byte(delim))
	return base.Buffer(val), err
}
func (r *StreamReader) Reset(rr io.Reader) {
	r.reader.Reset(rr)
}
func (r *StreamReader) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
func (r *StreamReader) WriteTo(ww io.Writer) (int64, error) {
	return r.reader.WriteTo(ww)
}
func (r *StreamReader) Seek(offset int64, whence int) (int64, error) {
	n, err := r.reader.Discard(int(offset))
	return int64(n), err
}
func (r *StreamReader) ReadRune() (rune, error) {
	val, _, err := r.reader.ReadRune()
	return val, err
}
func (r *StreamReader) ReadInt8() (uint8, error) {
	val, err := r.reader.ReadByte()
	return uint8(val), err
}
func (r *StreamReader) ReadInt16() (uint16, error) {
	buf, err := r.reader.Peek(2)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(2)
	return r.order.Uint16(buf), nil
}
func (r *StreamReader) ReadInt32() (uint32, error) {
	buf, err := r.reader.Peek(4)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(4)
	return r.order.Uint32(buf), nil
}
func (r *StreamReader) ReadInt64() (uint64, error) {
	buf, err := r.reader.Peek(8)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(8)
	return r.order.Uint64(buf), nil
}
