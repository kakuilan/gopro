// 指针
package main
import "fmt"

func main(){
    var a uintptr
    //var a *int
    b := 3
    //a = &b
    //a = (&b).(uintptr)
    c := (&b).(uintptr)
    fmt.Println(a,b,c)

}
