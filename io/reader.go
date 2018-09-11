package io

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/vlorc/lua-vm/base"
	"io"
	"strings"
)

type Reader interface {
	io.ReadSeeker
	Endian(bool)
	Reset(io.Reader)
	ReadStringAt(int) (string, error)
	ReadBufferAt(int) (base.Buffer, error)
	ReadLine() (string, error)
	ReadInt8() (int64, error)
	ReadInt16() (int64, error)
	ReadInt32() (int64, error)
	ReadInt64() (int64, error)
	ReadRune() (rune, error)
}

type StreamReader struct {
	reader *bufio.Reader
	order  binary.ByteOrder
}
type ReaderFactory struct{}

func (f ReaderFactory) New(b interface{}) Reader {
	switch r := b.(type) {
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
func (r *StreamReader) ReadInt8() (int64, error) {
	val, err := r.reader.ReadByte()
	return int64(val), err
}
func (r *StreamReader) ReadInt16() (int64, error) {
	buf, err := r.reader.Peek(2)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(2)
	return int64(r.order.Uint16(buf)), nil
}
func (r *StreamReader) ReadInt32() (int64, error) {
	buf, err := r.reader.Peek(4)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(4)
	return int64(r.order.Uint32(buf)), nil
}
func (r *StreamReader) ReadInt64() (int64, error) {
	buf, err := r.reader.Peek(8)
	if nil != err {
		return 0, err
	}
	r.reader.Discard(8)
	return int64(r.order.Uint64(buf)), nil
}
