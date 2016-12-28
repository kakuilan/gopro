// unsafe转换指针
package main
import "fmt"
import "unsafe"


func main(){
    u := uint32(32)
    i := int32(1)
    fmt.Println(&u, &i)

    p := &i
    p = (*int32)(unsafe.Pointer(&u))
    fmt.Println(p)



}

