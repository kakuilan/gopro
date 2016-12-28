// 要修改字符串,可先将其转换为[]rune 或 []byte,完成后再转换为string.无论哪种转换,都会重新分配内存,并复制字节数组
package main
import "fmt"

func main(){
    s := "abc"
    bs := []byte(s)
    bs[1] = 'B'
    println(string(bs))

    u := "电脑"
    us := []rune(u)
    us[1] = '话'
    println(string(us))

    fmt.Println(s,bs, u, us)
}
