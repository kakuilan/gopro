// defer 延迟函数,按照后进先出的顺序执行
package main
import "fmt"

func main(){
    s := "test"

    //将会打印 4 3 2 1 0
    for i := 0;i<5;i++{
	defer fmt.Println(i)
    }
    fmt.Println(s)
}
