// 用for循环遍历字符串时,也有byte和rune两种方式
package main
import "fmt"

func main(){
    s := "abc汉字"
    //byte
    for i:=0;i<len(s);i++{
	fmt.Printf("%c,", s[i])
    }

    fmt.Println()

    //rune
    for _,r := range s {
	fmt.Printf("%c,", r)
    }
    println()
}
