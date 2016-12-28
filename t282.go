// 支持用两个索引号返回子串
package main
import "fmt"

func main(){
    s := "Hello, World!"
    s1 := s[:5] //Hello
    s2 := s[7:] //World!
    s3 := s[1:5] //ello
    fmt.Println(s1, s2, s3)

}
