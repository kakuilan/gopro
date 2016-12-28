// 合法的unsafe.Pointer被当做普通指针对待
package main
import (
		"unsafe"
		"time"
		)

type data struct {
	x [1024 * 1024] byte
}

func test() unsafe.Pointer {
	p := &data{}
	return unsafe.Pointer(p)
}

func main(){
	const N = 10000
	cache := new([N]unsafe.Pointer)

	for i:=0;i<N;i++{
		cache[i] = test()
		time.Sleep(time.Millisecond)
	
	}


}
