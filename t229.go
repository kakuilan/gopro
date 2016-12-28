// 函数的变参
package main
import "fmt"

func f1(arg ...int){
    fmt.Println("f1函数的参数:", arg)
    f2(arg...) //按原样传递
    f3(arg[:3]...) //只传递部分
}

func f2(arg ...int){
    fmt.Println("f2函数的参数:", arg)
}

func f3(arg ...int) {
    fmt.Println("f3函数的参数:", arg)
}

func main(){
    f1(1,2,3,4,5,6,7)
}
