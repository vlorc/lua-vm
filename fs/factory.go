package fs

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type RelativeFileFactory struct {
	parent FileSystem
	root   string
}

func __parse(root, src string) (string, error) {
	dst := filepath.Join(root, src)
	dst, err := filepath.Abs(dst)
	if nil != err {
		return "", err
	}
	if !strings.HasPrefix(dst, root) {
		return "", fmt.Errorf("Can't open file %s to %s", src, dst)
	}
	return dst, nil
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

func (f *RelativeFileFactory) Open(file string, args ...int) (FileDriver, error) {
	dst, err := __parse(f.root, file)
	if nil != err {
		return nil, err
	}
	return f.parent.Open(dst)
}

func (f *RelativeFileFactory) Remove(file string) error {
	dst, err := __parse(f.root, file)
	if nil == err {
		err = f.parent.Remove(dst)
	}
	return err
}
