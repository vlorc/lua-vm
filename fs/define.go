package fs

import (
	"io"
	"os"
)

type FileDriver interface {
	io.ReadWriteCloser
	io.Seeker
	Stat() (os.FileInfo, error)
}

type FileSystem interface {
	Open(file string, args ...int) (FileDriver, error)
	Remove(file string) error
	/*Remove(src,dst string) error
	MkDir(file string) error
	MkTmp(file string) (string,error)*/
}
