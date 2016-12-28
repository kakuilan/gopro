// 使用索引号访问字符 byte
package main
import "fmt"

func main(){
    s := "abc"
    fmt.Println(s[0]=='\x61', s[1]=='b', s[2]==0x63)

    //使用`定义不做转义处理的原始字符串,支持跨行
    s2 := `
    b\r\nx00
    c`
    fmt.Println(s2)

}
