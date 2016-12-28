// 字符串替换
package main
import "fmt"

func main(){
    s := "qwe writer123423 "
    r := []rune(s)
    copy(r[4:4+3], []rune("abc"))
    fmt.Println(s)
    fmt.Println(string(r))
}
