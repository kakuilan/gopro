// 统计字符串里的字符数量和字符串的字节数
package main
import (
    "fmt"
    "unicode/utf8"
)

func main(){
    str := "asSASA ddd dsjkdsjs dk"
    strLen := len([]byte(str))
    strByt := utf8.RuneCount([]byte(str))
    fmt.Println("字符数量:", strLen)
    fmt.Println("字节数:", strByt)
}
