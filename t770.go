//类型断言
package main
import "fmt"

func main() {
    var i interface{} = 99
    var s interface{} = []string{"left","right" }

    j := i.(int) //非安全类型断言
    fmt.Printf("%T->%d\n", j, j)

    //安全类型断言
    if i,ok := i.(int);ok {
        fmt.Printf("%T->%d\n", i, j) //i是一个int类型的影子变量
    }

    if s,ok := s.([]string);ok {
        fmt.Printf("$T->%q\n", s, s) //s是一个[]string类型的影子变量
    }

}
