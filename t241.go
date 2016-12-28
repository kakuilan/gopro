// 指针
package main
import "fmt"

func main(){
    var p *int
    fmt.Println(p)

    var i int
    p = &i
    *p = 8 //修改i的值
    fmt.Println(p,i)

}
