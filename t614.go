//命名返回值参数
package main
import "fmt"

func main(){
	sum,sub := f3(3,6)
	fmt.Println(sum,sub)
}

func f3(a,b int) (sum,sub int) {
	sum = a+b
	sub = a-b
	return
}
