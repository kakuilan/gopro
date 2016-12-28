package main
import "fmt"

//打印输入
func main(){
    var i = 0
    var s = ""
    fmt.Scanf("%s is %d", &s, &i)
    fmt.Printf("%s is %d", s, i)
}
