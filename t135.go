//函数变参

package main
import "fmt"

func sum(aregs ...int) int {
    s := 0
    for _,number := range aregs {
	s += number
    }
    return s
}

func main(){
    total := sum(1,2,3,4)
    fmt.Println(total)
    slice := []int {1,2,3,4,5,6,7,8,9} //定义一个切片
    //将切片传sum时,要用...展开,否则将作为一个参数处理
    //等价于 sum(1,2,3,4,5,6,7,8,9)
    total = sum(slice...)
    fmt.Println(total)

}
