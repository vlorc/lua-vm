package fs

import (
	"io/fs"
	"os"
	"path/filepath"
)

type VirtualFile struct {
	fd fs.File
}

func (v *VirtualFile) Read(p []byte) (n int, err error) {
	return v.fd.Read(p)
}

func (v *VirtualFile) Write(p []byte) (n int, err error) {
	return 0, ErrMethodNotSupport
}

func (v *VirtualFile) Close() error {
	return v.fd.Close()
}

func (v *VirtualFile) Seek(offset int64, whence int) (int64, error) {
	return 0, ErrMethodNotSupport
}

func (v *VirtualFile) Stat() (os.FileInfo, error) {
	return v.fd.Stat()
}

type VirtualFileFactory struct {
	parent FileSystem
	fileSys fs.FS
}

var _ FileDriver = &VirtualFile{}

func (v VirtualFileFactory) Open(file string, args ...int) (FileDriver, error) {
	f, err := v.fileSys.Open(file)
	if nil != err {
		return nil, err
	}

	return &VirtualFile{fd: f}, nil
}

func (VirtualFileFactory) Rename(src, dst string) error {
	return ErrMethodNotSupport
}

func (VirtualFileFactory) Remove(file string) error {
	return ErrMethodNotSupport
}

func (v VirtualFileFactory) Exist(file string) bool {
	f, err := v.fileSys.Open(file)
	if nil != err {
		return false
	}
	_ = f.Close()
	return true
}

func (VirtualFileFactory) Mkdir(file string, mode int) error {
	return ErrMethodNotSupport
}

func (VirtualFileFactory) Walk(root string, callback filepath.WalkFunc) error {
	return ErrMethodNotSupport
}
