// 引用类型
package main
import "fmt"

func main(){
    a := []int{0,0,0} //提供初始化表达式
    a[1] = 10

    b := make([]int, 3) //make slice
    b[1] = 10

    c := new([]int)
    //c[1] = 10 //将发生错误 invalid operation: c[1] (index of type *[]int)

    fmt.Println(a,b,c)

}
