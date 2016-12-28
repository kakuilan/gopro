// 指针构成的"循环引用"加上 runtime.SetFinalizer 会导致内存泄露
package main
import (
		"runtime"
		"time"
		"fmt"
		)

type Data struct {
	d [1024 * 1024] byte
	o *Data
}

func test(){
	var a,b Data
	//循环引用
	a.o = &b
	b.o = &a

	runtime.SetFinalizer(&a, func(d *Data){fmt.Printf("a %p final.\n", d)})
	runtime.SetFinalizer(&b, func(d *Data){fmt.Printf("b %p final.\n", d)})
}

func main(){
	for{
		test()
		time.Sleep(time.Millisecond)
	
	}

}


