// 类型转换
package main
import "fmt"

func main(){
    mystring := "Hello this is a string"
    //转换到byte slice,每个byte保存字符串对应字节的整数值,UTF-8编码
    byteslice := []byte(mystring)

    //转换到rune slice,每个rune保存Unicode编码的指针
    runeslice := []rune(mystring)

    //从字节或整形的slice到string
    b := []byte{'h','e','l','l','o'}
    s := string(b)
    i := []rune{257,1024, 65}
    r := string(i)

    fmt.Println(byteslice)
    fmt.Println(runeslice)
    fmt.Println(b)
    fmt.Println(s)
    fmt.Println(i)
    fmt.Println(r)


}
