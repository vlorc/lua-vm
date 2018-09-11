package fs

import "io"

type FileDriver interface {
	io.ReadWriteCloser
	io.Seeker
}

type FileSystem interface {
	Open(file string) (FileDriver,error)
	/*Create(file string,mode int) (File,error)
	Remove(src,dst string) error
	Delete(file string) error
	MkDir(file string) error
	MkTmp(file string) (string,error)*/
}