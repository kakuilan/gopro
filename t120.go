//变量定义
package main
import "fmt"

func main(){
    var b bool
    var n int
    var i int = 3
    var (//多个变量的定义
	aa int = 3
	str string
    )

    //定义多个变量并赋值
    var i1,i2,i3 int = 1,2,3
    var strName = "张三"
    strSex := "男"

    fmt.Println("n=", n)
    fmt.Println("b=", b)
    fmt.Println("i=", i)
    fmt.Println("aa=", aa)
    fmt.Println("str=", str)
    fmt.Println("i1=", i1, ",i2=", i2, ",i3=", i3)
    fmt.Println("strName=", strName)
    fmt.Println("strSex=", strSex)
}
