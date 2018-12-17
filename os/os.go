package os

import (
	"runtime"
)

type OSFactory struct{}

func (OSFactory) Version() string {
	return runtime.Version()
}

func (OSFactory) Name() string {
	return runtime.GOOS
}

func (OSFactory) Arch() string {
	return runtime.GOARCH
}
