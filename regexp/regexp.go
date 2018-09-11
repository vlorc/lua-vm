package regexp

import (
	"github.com/vlorc/lua-vm/base"
	"io"
	"regexp"
)

type RegexpFactory struct{}

func (f RegexpFactory) New(expr string) (*regexp.Regexp, error) {
	return f.Compile(expr)
}

func (RegexpFactory) Compile(expr string) (*regexp.Regexp, error) {
	return regexp.Compile(expr)
}

func (RegexpFactory) CompilePOSIX(expr string) (*regexp.Regexp, error) {
	return regexp.CompilePOSIX(expr)
}

func (RegexpFactory) Match(pattern string, b base.Buffer) (bool, error) {
	return regexp.Match(pattern, b)
}

func (RegexpFactory) MatchReader(pattern string, r io.RuneReader) (bool, error) {
	return regexp.MatchReader(pattern, r)
}

func (RegexpFactory) MatchString(pattern, s string) (bool, error) {
	return regexp.MatchString(pattern, s)
}

func (RegexpFactory) MustCompile(str string) *regexp.Regexp {
	return regexp.MustCompile(str)
}

func (RegexpFactory) MustCompilePOSIX(str string) *regexp.Regexp {
	return regexp.MustCompilePOSIX(str)
}

func (RegexpFactory) QuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}
