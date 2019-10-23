package log

import (
	"bytes"
	"runtime"
	"unsafe"
)

// Goid get gourtine id
func Goid() string {
	var buf [32]byte
	s := buf[10:runtime.Stack(buf[:], false)]
	s = s[:bytes.IndexByte(s, ' ')]
	return *(*string)(unsafe.Pointer(&s))
}
