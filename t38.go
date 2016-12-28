package main
import "fmt"

//定义新类型Integer,并为它增加新方法Less()
type Integer int
func (a Integer) Less(b Integer) bool{
    return a < b
}

func main(){
    var a Integer = 1
    if a.Less(2){
	fmt.Println("a Less 2")
    }else{
	fmt.Println("a not less 2")
    }
}
