package main

import (
	"runtime"
	"strconv"
	"unsafe"
)

// Stub out main instead of requiring host to pass one.
func main() {}

//export work
func work(i uint32) uint32 {
	// Churn the GC to make sure unreferenced memory is free.
	for i := 0; i < 1000; i++ {
		a := "hello" + strconv.Itoa(i)
		if len(a) < 0 {
			break
		}
		runtime.GC()
	}

	// Add conditional logic to make sure compiler doesn't remove.
	var s string
	if i == 0 {
		s = "pandabear"
	} else {
		s = "polarbear"
	}

	buf := []byte(s)
	ptr := &buf[0]
	return uint32(uintptr(unsafe.Pointer(ptr)))
}
