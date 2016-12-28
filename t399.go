// 通过cgo,可在GO和C/C++代码间相互调用
package main

/*
	#include <stdio.h>
	#include <stdlib.h>

	void hello(){
		printf("Hello, World! from c.\n")
	}
 
 */

import "C"

func main(){
	C.hello()
}
