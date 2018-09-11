package fs

import (
	"errors"
	"fmt"
	"io"
	"github.com/vlorc/lua-vm/net/http"
)

type HttpFile struct {
	fd io.ReadCloser
}

type HttpFileFactory struct{
	root string
	driver *http.HTTPFactory
}

func(f *HttpFileFactory)Open(file string) (FileDriver,error) {
	resp,err := f.driver.Get(f.root + file)
	if nil != err {
		return nil,err
	}
	if 200 != resp.StatusCode {
		return nil,fmt.Errorf("Can't open http file code: %d", resp.StatusCode)
	}
	return &HttpFile{fd: resp.Body},nil
}

func(f *HttpFile)Write(b []byte)(int,error) {
	return 0,errors.New("Can't support write method")
}

func(f *HttpFile)Read(b []byte)(int,error) {
	return f.fd.Read(b)
}

func(f *HttpFile)Close()(error) {
	return f.fd.Close()
}

func(f *HttpFile)Seek(offset int64, whence int) (int64, error){
	return 0,errors.New("Can't support seek method")
}