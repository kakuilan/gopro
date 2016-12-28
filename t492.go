// 中文类型别名
package main
import "fmt"

type (
		小数 float32
		文本 string
		)

func main(){
	var f 小数 = 3.14
	var str 文本 = "中国"
	fmt.Println(f)
	fmt.Println(str)
}


