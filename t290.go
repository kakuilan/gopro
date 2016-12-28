// unsafe.Pointer和任意类型指针间进行转换
package main
import (
    "fmt"
    "unsafe"
)

func main(){
    x := 0x12345678

    p := unsafe.Pointer(&x) // *int -> Pointer
    n := (*[4]byte)(p) // Pointer -> *[4]byte

    for i := 0;i<len(n);i++{
	//%X把十进制转换16进制
	fmt.Printf("%X \n", n[i])
    }

    fmt.Println(x,p, n)


}
