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
	v, err := __parse(f.root, file)
	if nil != err {
		return nil, err
	}
	return f.parent.Open(v, args...)
}

func (f *RelativeFileFactory) Remove(file string) error {
	v, err := __parse(f.root, file)
	if nil == err {
		err = f.parent.Remove(v)
	}
	return err
}

func (f *RelativeFileFactory) Rename(src, dst string) error {
	s, e := __parse(f.root, src)
	if nil != e {
		return e
	}
	d, e := __parse(f.root, dst)
	if nil == e {
		e = os.Rename(s, d)
	}
	return e
}

func (f *RelativeFileFactory) Exist(file string) bool {
	v, err := __parse(f.root, file)
	if nil == err {
		return f.parent.Exist(v)
	}
	return false
}

func (f *RelativeFileFactory) Mkdir(file string, mode int) error {
	v, err := __parse(f.root, file)
	if nil == err {
		err = f.parent.Mkdir(v, mode)
	}
	return err
}

func (f *RelativeFileFactory) Walk(root string, callback filepath.WalkFunc) error {
	v, err := __parse(f.root, root)
	if nil == err {
		err = f.parent.Walk(v, callback)
	}
	return err
}
