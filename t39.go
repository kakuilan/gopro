package main
import "fmt"

//定义新类型Integer,并为它增加新方法Add()
type Integer int
func (a *Integer) Add(b Integer){
    *a += b
}

func main(){
    var a Integer = 1
    a.Add(2)
    fmt.Println("a = ", a)

}
