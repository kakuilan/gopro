//修改切片
package main
import "fmt"

func main() {
    s := []string {"A","B","C","D","E","F","G"}
    t := []string {"K","L","M","N"}
    u := []string {"m","n","o","p","q","r"}

    s = append(s, "h", "i", "j") //添加单一的值
    s = append(s, t...) //添加切片中的所有值
    s = append(s, u[2:5]...) //添加一个子切片
    b := []byte{'U','V'} 
    letters := "WXY"
    b = append(b, letters...) //将一个字符串字节添加进一个字节切片中
    fmt.Printf("%v\n%s\n", s, b)




}
