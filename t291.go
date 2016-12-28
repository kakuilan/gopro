// 返回局部变量指针是安全的,编译器会根据需要将其分配在GC Heap上
package main
import "fmt"

func test() *int {
    x := 100
    return &x
}

func main(){
    var a *int
    a = test()
    fmt.Println(a)
}
