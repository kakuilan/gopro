// 将Pointer转换为uintptr,可变相实现指针运算
package main
import "fmt"
import "unsafe"

func main(){
    d := struct {
	s string
	x int
    }{"abc", 100}
    fmt.Println(d)

    p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
    fmt.Println(unsafe.Offsetof(d.s), unsafe.Offsetof(d.x))
    p += unsafe.Offsetof(d.x) // uintptr + offset
    p2 := unsafe.Pointer(p) // uintptr -> Pointer
    px := (*int)(p2) // Pointer -> *int
    *px = 200 // d.x = 200

    fmt.Printf("%#v\n", d)


}
