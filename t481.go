//算术运算
package main
import "fmt"

func main(){
	var a,b int = 7,163
	var c,d float32 = 2.5,6.3
	var ch1,ch2 rune = '中','国'
	fmt.Println(string(ch1) + string(ch2)) //加法
	fmt.Println(d -c) //减法
	fmt.Println(c * d) //乘法
	fmt.Println(b / a) //除法
	fmt.Println(b % a) //取余


}
