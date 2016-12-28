//空interface可存储任意类型的值
package main
import "fmt"

func main(){
	//定义a为空接口
	var a interface{}
	var i int = 5
	s := "Hello World"
	//a可以存储任意类型的数值
	a = i
	a = s
	fmt.Println(a)


}



