//go:build !go1.7
// +build !go1.7

package is

import (
	"regexp"
)

var reIsSourceFile = regexp.MustCompile("is(-before-1.7)?\\.go$")

func (is *I) callerinfo() (path string, line int, ok bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}
