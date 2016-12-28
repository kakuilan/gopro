// 初始化值以","分隔.可以分多行,但最后一行必须以","或"}"结尾
package main
import "fmt"

func main(){
    //错误的
    //a := []int{
    //	1,
    //  2 //这里出错
    //}

    a1 := []int{
	1,
	2, //ok
    }

    a2 := []int{
	1,
	2 } //ok
    fmt.Println(a1,a2)
}
