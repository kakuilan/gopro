// 函数作为值
package main
import "fmt"

func main(){
    a := func(){ //定义一个匿名函数,并赋值给a
	println("Hello")
    }
    a() //调用函数
    fmt.Printf("%T\n", a) //打印a的类型

    //使用map的函数作为值
    var xs = map[int]func() int{
	1: func() int {return 10},
	2: func() int {return 20},
	3: func() int {return 30}, //这个逗号必须有
    }
    r := xs[1]()
    fmt.Println(r, xs)

}
