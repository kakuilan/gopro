// 字符串转换函数
package main

/*
#include <stdio.h>
#include <stdlib.h>

void test(char *s){
	printf("%s\n", s);
}

char* cstr(){
	return "abcde";
}
 
 
*/

import "C"
import "fmt"

func main(){
	s := "Hello, World!"
	cs := C.CString(s) //该函数在C heap分配内存,需要调用free释放
	defer C.free(unsafe.Pointer(cs))

	C.test(cs)
	cs = C.cstr()

	fmt.Println(C.GoString(cs))
	fmt.Println(C.GoString(cs, 2))
	fmt.Println(C.GoBytes(unsafe.Pointer(cs), 2))

}
