// 定义自己的类型
package main
import "fmt"

type foo int

type NameAge struct {
    name string //小写,不导出
    age int
}

func main(){
    var f foo
    f = 4

    a := new(NameAge)
    a.name = "pete"
    a.age = 32

    fmt.Println(f, a)


}
