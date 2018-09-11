package tcp

import (
	"context"
	"time"
)

func Context(ctx context.Context) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.context = ctx
	}
}

func Timeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.resolveTimeout = timeout
		opt.listenTimeout = timeout
		opt.connectTimeout = timeout
		opt.readTimeout = timeout
		opt.writeTimeout = timeout
	}
}
func ConnectTimeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.connectTimeout = timeout
	}
}
func ListenTimeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.listenTimeout = timeout
	}
}
func ResolveTimeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.resolveTimeout = timeout
	}
}
func ReadTimeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.readTimeout = timeout
	}
}
func WriteTimeout(timeout time.Duration) func(*TCPFactory) {
	return func(opt *TCPFactory) {
		opt.writeTimeout = timeout
	}
}

func (o *TCPFactory) __context(timeout time.Duration, args []int) (context.Context, context.CancelFunc) {
	if len(args) > 1 {
		timeout = time.Duration(args[1]) * time.Microsecond
	}
	if timeout > 0 {
		return context.WithTimeout(o.context, timeout)
	}
	return o.context, func() {}
}
