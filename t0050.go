//cgo例子,引用C库
package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

func main() {
	Print("Hello world\n")
}
