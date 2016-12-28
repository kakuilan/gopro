package main
import "fmt"

//输出输入到字符串
func main(){
    var i =0
    var s = ""
    var t = "Answer is 54"
    fmt.Sscanf(t, "%2s is %d", &s, &i)
    fmt.Println(s, i)
}
