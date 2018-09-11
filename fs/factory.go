package fs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type NativeFile struct {
}

type NativeFileFactory struct{}

func (NativeFileFactory) Open(file string) (FileDriver, error) {
	fd, err := os.Open(file)
	return fd, err
}

type RelativeFileFactory struct {
	parent FileSystem
	root   string
}

func NewRelativeFileFactory(root string, parent FileSystem) FileSystem {
	if !filepath.IsAbs(root) {
		app, _ := exec.LookPath(os.Args[0])
		path, _ := filepath.Abs(app)
		root = filepath.Join(filepath.Dir(path), root)
	}
	return &RelativeFileFactory{
		parent: parent,
		root:   root,
	}
}

func (f *RelativeFileFactory) Open(src string) (FileDriver, error) {
	dst := filepath.Join(f.root, src)
	dst, err := filepath.Abs(dst)
	if nil != err {
		return nil, err
	}
	if !strings.HasPrefix(dst, f.root) {
		return nil, fmt.Errorf("Can't open file %s to %s", src, dst)
	}
	return f.parent.Open(dst)
}
