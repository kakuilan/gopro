// uintprt被GC当做普通整数对象,它不能阻止所"引用"对象被回收
package main
import "fmt"
import "unsafe"
import "time"

type data struct {
	x [1024 * 1024]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}

func main(){
	const N = 10000
	cache := new([N]uintptr)
	fmt.Println()
	for i:=0;i<N;i++{
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
}


