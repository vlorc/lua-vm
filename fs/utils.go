package fs

import (
	"bytes"
	"github.com/vlorc/lua-vm/base"
	"unsafe"
)

type FileUtilsFactory struct {
	FileSystem
}

func (f FileUtilsFactory) ReadString(file string, args ...int) (string, error) {
	buf, err := f.__readBuffer(file, args...)
	if nil != err {
		return "", nil
	}
	return *(*string)(unsafe.Pointer(&buf)), nil
}

func (f FileUtilsFactory) ReadBuffer(file string, args ...int) (base.Buffer, error) {
	buf, err := f.__readBuffer(file, args...)
	if nil != err {
		return nil, nil
	}
	return base.Buffer(buf), nil
}

func (f FileUtilsFactory) __readBuffer(file string, args ...int) ([]byte, error) {
	fd, err := f.Open(file)
	if nil != err {
		return nil, err
	}
	defer fd.Close()

	var n int64 = bytes.MinRead
	if fi, err := fd.Stat(); err == nil {
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}
	}
	var buf bytes.Buffer
	buf.Grow(int(n))
	_, err = buf.ReadFrom(fd)
	return buf.Bytes(), err
}

func (f FileUtilsFactory) WriteString(file string, str string, args ...int) (int, error) {
	fd, err := f.Open(file, 0, 0666)
	if nil != err {
		return 0, err
	}
	defer fd.Close()
	return fd.Write(*(*[]byte)(unsafe.Pointer(&str)))
}

func (f FileUtilsFactory) WriteBuffer(file string, buf base.Buffer, args ...int) (int, error) {
	fd, err := f.Open(file, 0, 0666)
	if nil != err {
		return 0, err
	}
	defer fd.Close()
	return fd.Write(buf)
}
