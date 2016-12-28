//字符串的定义和初始化
package main
import "fmt"

func main(){
	var str string
	str = "Hello World!"
	fmt.Println(str)
	fmt.Println("字符串长度：", len(str))
	fmt.Println("第一个字符：ASCII码 ", str[0], " 实际字符：", string(str[0]))

}
