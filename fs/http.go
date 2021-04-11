package fs

import (
	"fmt"
	vmhttp "github.com/vlorc/lua-vm/net/http"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type HttpFile struct {
	fd     io.ReadCloser
	url    string
	length string
	modify string
}

type HttpFileInfo struct {
	name   string
	length int64
	modify time.Time
}

type HttpFileFactory struct {
	root   string
	driver *vmhttp.HTTPFactory
}

func (f *HttpFileFactory) Open(file string, args ...int) (FileDriver, error) {
	resp, err := f.driver.Get(f.root + file)
	if nil != err {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("Can't open http file code: %d", resp.StatusCode)
	}
	return &HttpFile{
		fd:     resp.Body,
		url:    file,
		length: resp.Header.Get("Content-Length"),
		modify: resp.Header.Get("Last-Modified"),
	}, nil
}

func (f *HttpFileFactory) Remove(file string) error {
	resp, err := f.driver.Delete(f.root + file)
	if nil == err {
		defer resp.Body.Close()
		if 200 != resp.StatusCode {
			err = fmt.Errorf("Can't open http file code: %d", resp.StatusCode)
		}
	}
	return err
}

func (*HttpFileFactory) Rename(src, dst string) error {
	return ErrMethodNotSupport
}

func (f *HttpFileFactory) Exist(file string) bool {
	resp, err := f.driver.Head(f.root + file)
	if nil == err {
		defer resp.Body.Close()
		if 200 != resp.StatusCode {
			err = fmt.Errorf("Can't exist http file code: %d", resp.StatusCode)
		}
	}
	return true
}

func (*HttpFileFactory) Mkdir(string, int) error {
	return ErrMethodNotSupport
}

func (*HttpFileFactory) Walk(root string, callback filepath.WalkFunc) error {
	return ErrMethodNotSupport
}

func (f *HttpFile) Write(b []byte) (int, error) {
	return 0, ErrMethodNotSupport
}

func (f *HttpFile) Read(b []byte) (int, error) {
	return f.fd.Read(b)
}

func (f *HttpFile) Close() error {
	return f.fd.Close()
}

func (f *HttpFile) Seek(offset int64, whence int) (int64, error) {
	return 0, ErrMethodNotSupport
}

func (f *HttpFile) Stat() (os.FileInfo, error) {
	_, filename := filepath.Split(f.url)
	length, _ := strconv.Atoi(f.length)
	last, _ := time.Parse(http.TimeFormat, f.modify)
	return &HttpFileInfo{
		name:   filename,
		length: int64(length),
		modify: last,
	}, nil
}

func (f *HttpFileInfo) Name() string {
	return f.name
}
func (f *HttpFileInfo) Size() int64 {
	return f.length
}
func (f *HttpFileInfo) Mode() os.FileMode {
	return 0
}
func (f *HttpFileInfo) ModTime() time.Time {
	return f.modify
}
func (f *HttpFileInfo) IsDir() bool {
	return false
}
func (f *HttpFileInfo) Sys() interface{} {
	return nil
}
