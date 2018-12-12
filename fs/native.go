package fs

import (
	"os"
	"path/filepath"
)

type NativeFile struct{}

type NativeFileFactory struct{}

func (NativeFileFactory) Open(file string, args ...int) (FileDriver, error) {
	flag := os.O_RDWR
	mode := 0666
	if len(args) > 0 {
		if 0 == args[0] {
			flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
		} else {
			flag = args[0]
		}
		if len(args) > 1 {
			mode = args[0]
		}
	}
	fd, err := os.OpenFile(file, flag, os.FileMode(mode))
	return fd, err
}

func (NativeFileFactory) Rename(src, dst string) error {
	return os.Rename(src, dst)
}

func (NativeFileFactory) Remove(file string) error {
	return os.Remove(file)
}

func (NativeFileFactory) Exist(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (NativeFileFactory) Mkdir(file string, mode int) error {
	return os.MkdirAll(file, os.FileMode(mode))
}

func (NativeFileFactory) Walk(root string, callback filepath.WalkFunc) error {
	return filepath.Walk(root, callback)
}
