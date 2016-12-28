// 指针对象成员的unsafe.Pointer,同样能确保对象不被回收
package main
import (
		"unsafe"
		"time"
		)

type data struct {
	x [1024 * 1024] byte
	y int
}

func test() unsafe.Pointer {
	d := data{}
	//指向对象成员y的unsafe.Pointer
	return unsafe.Pointer(&d.y)
}

func main(){
	const N = 10000
	cache := new([N]unsafe.Pointer)
	
	for i:=0;i<N;i++{
		cache[i] = test()
		time.Sleep(time.Millisecond)
	
	}
}

