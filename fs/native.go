package fs

import "os"

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

func (NativeFileFactory) Remove(file string) error {
	return os.Remove(file)
}
