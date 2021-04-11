package fs

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrMethodNotSupport = errors.New("Can't support method")
)

type FileDriver interface {
	io.ReadWriteCloser
	io.Seeker
	Stat() (os.FileInfo, error)
}

type FileSystem interface {
	Open(file string, args ...int) (FileDriver, error)
	Remove(file string) error
	Rename(src, dst string) error
	Exist(file string) bool
	Mkdir(file string, mode int) error
	Walk(root string, callback filepath.WalkFunc) error
}
