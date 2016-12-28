// 初始化复合对象,必须使用类型标签,且左大括号必须在类型尾部
package main
import "fmt"

func main(){
    //var a struct {x int} = {100} 错误
    //var b []int = {1,2,3} 错误
    //c := struct {x int; y sting}
    //{
    //}

    //正确的
    var a = struct{x int}{100}
    var b = []int{1,2,3}
    fmt.Println(a,b)

}
