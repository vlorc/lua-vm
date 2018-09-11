package io

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/vlorc/lua-vm/base"
	"io"
)

type WriterFactory struct{}

type Writer interface {
	io.Writer
	Endian(bool)
	Reset(io.Writer)
	Flush() error
	Len() int
	Size() int
	ReadFrom(io.Reader) (int64, error)
	WriteString(string) (int, error)
	WriteLine(string) (int, error)
	WriteInt8(int) (int, error)
	WriteInt16(int) (int, error)
	WriteInt32(int) (int, error)
	WriteInt64(int) (int, error)
}

type StreamWriter struct {
	writer *bufio.Writer
	order  binary.ByteOrder
}

func (f WriterFactory) New(b interface{}) Writer {
	switch r := b.(type) {
	case []byte:
		return f.FormBuffer(base.Buffer(r))
	case base.Buffer:
		return f.FormBuffer(base.Buffer(r))
	case string:
		return f.FormString(r)
	case io.Writer:
		return f.FormStream(r)
	}
	return nil
}

func (WriterFactory) FormStream(w io.Writer) Writer {
	return &StreamWriter{
		writer: bufio.NewWriter(w),
		order:  binary.LittleEndian,
	}
}

func (f WriterFactory) FormSize(size int) Writer {
	return f.FormStream(bytes.NewBuffer(make([]byte, size)))
}
func (f WriterFactory) FormString(str string) Writer {
	return f.FormStream(bytes.NewBufferString(str))
}
func (f WriterFactory) FormBuffer(buf base.Buffer) Writer {
	return f.FormStream(bytes.NewBuffer(buf))
}
func (w *StreamWriter) Endian(big bool) {
	if big {
		w.order = binary.BigEndian
	} else {
		w.order = binary.LittleEndian
	}
	w.writer.Size()
}
func (w *StreamWriter) Len() int {
	return w.writer.Buffered()
}
func (w *StreamWriter) Size() int {
	return w.writer.Size()
}
func (w *StreamWriter) Reset(ww io.Writer) {
	w.writer.Reset(ww)
}
func (w *StreamWriter) Flush() error {
	return w.writer.Flush()
}
func (w *StreamWriter) WriteLine(str string) (n int, err error) {
	if n, err = w.writer.WriteString(str); nil != err {
		return 0, err
	}
	return n + 1, w.writer.WriteByte(byte('\n'))
}
func (w *StreamWriter) WriteString(str string) (n int, err error) {
	return w.writer.WriteString(str)
}
func (w *StreamWriter) Write(p []byte) (n int, err error) {
	return w.writer.Write(p)
}
func (w *StreamWriter) ReadFrom(r io.Reader) (int64, error) {
	return w.writer.ReadFrom(r)
}
func (w *StreamWriter) WriteRune(v rune) (int, error) {
	return w.writer.WriteRune(v)
}
func (w *StreamWriter) WriteInt8(v int) (int, error) {
	return 1, w.writer.WriteByte(byte(v & 0xff))
}
func (w *StreamWriter) WriteInt16(v int) (int, error) {
	var buf [2]byte
	w.order.PutUint16(buf[:], uint16(v&0xffff))
	return w.writer.Write(buf[:])
}
func (w *StreamWriter) WriteInt32(v int) (int, error) {
	var buf [4]byte
	w.order.PutUint32(buf[:], uint32(v&0xffffffff))
	return w.writer.Write(buf[:])
}
func (w *StreamWriter) WriteInt64(v int) (int, error) {
	var buf [8]byte
	w.order.PutUint64(buf[:], uint64(v))
	return w.writer.Write(buf[:])
}
