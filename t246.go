// 自定义类型转换
package main
import "fmt"

type foo struct {int} //匿名字段
type bar foo //bar是foo的别名

func main(){
    var b bar = bar{1} //声明b为bar类型
    //var f foo = b //赋值b到f, 将会引起错误, 不能使用b(类型bar)作为类型foo赋值
    //可以通过转换来修复
    var f foo = foo(b)
    fmt.Println(b,f)
}
