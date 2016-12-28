// func 传值与传指针
// func传参数时,实际上是传了该值的一份copy
package main
import "fmt"

func add1(a int) int {
	a = a+1 //改变a的值
	return a //返回一个新值
}

func main(){
	x := 3
	fmt.Println("x = ", x)

	x1 := add1(x)
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x) //x没有变化


}






