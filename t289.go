// 指针, &取变量地址, *透过指针访问目标对象
package main
import "fmt"

func main(){
    type data struct {a int}

    var d = data{1234}
    var p *data

    p = &d
    fmt.Printf("%p, %v\n", p, p.a) //直接用指针访问目标对象成员,无须转换

    //不能对指针做加减等运算
    x := 1234
    p2 := &x
    p2 ++ 


}
