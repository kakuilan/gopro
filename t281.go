// 连接跨行字符串时,"+"必须在上一行末尾,否则编译错误
package main
import "fmt"

func main(){
    s := "Hello, "+
	 "World!"
    fmt.Println(s)

}
