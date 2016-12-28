// 函数返回函数
package main
import "fmt"

func main(){
    p2 := plusTwo()
    p4 := plusX(4)
    fmt.Println(p2(2))
    fmt.Println(p4(5))
}

func plusTwo() func(int) int {
    return func(x int) int {return x+2}
}

func plusX(x int) func(int) int {
    return func(y int) int {return x+y}
}
