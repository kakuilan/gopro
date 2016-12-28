// defer 修改函数返回值
package main
import "fmt"

func f()(ret int){
    defer func(){
	ret++
    }()
    return 1
}
func main(){
    i := f() //返回值2
    fmt.Println(i)
}
