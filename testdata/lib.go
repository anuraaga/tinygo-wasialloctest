package main

import "unsafe"

//export malloc
func malloc(size uintptr) unsafe.Pointer

//export free
func free(ptr unsafe.Pointer)

//export get_buf
func getBuf() unsafe.Pointer {
	// To make clear that Go should have no reason to track this buffer with GC, we don't
	// cast to a Go type and let the memory be populated in the host. This is similar to
	// how memory would work if it is a file compiled in a different language such as C
	// that is calling malloc.
	return malloc(4)
}

//export release_buf
func releaseBuf(ptr unsafe.Pointer) {
	free(ptr)
}
