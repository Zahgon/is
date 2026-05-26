//go:build go1.7
// +build go1.7

package is

import (
	"regexp"
)

// Helper marks the calling function as a test helper function.
// When printing file and line information, that function will be skipped.
//
// Available with Go 1.7 and later.
func (is *I) Helper() { _ = "STUB: not implemented"; return }

// callerName gives the function name (qualified with a package path)
// for the caller after skip frames (where 0 means the current function).
func callerName(skip int) string {
	_ = "STUB: not implemented"
	// Make room for the skip PC.
	return ""
}

// skip + runtime.Callers + callerName

// The maximum number of stack frames to go through when skipping helper functions for
// the purpose of decorating log messages.
const maxStackLen = 50

var reIsSourceFile = regexp.MustCompile(`is(-1.7)?\.go$`)

func (is *I) callerinfo() (path string, line int, ok bool) {
	_ = "STUB: not implemented"
	return "",

		// Skip two extra frames to account for this function
		// and runtime.Callers itself.
		0, false
}

// Frame is inside a helper function.

// If no "non-helper" frame is found, the first non is frame is returned.
