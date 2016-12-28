// func传指针,需要将参数x所在的地址&x传入函数,并将参数的类型由int改为*int指针类型
// 此时,参数仍然是按copy传递的,只是copy的是一个指针.
package main
import "fmt"

func add1(a *int) int {
	*a = *a+1
	return *a
}

func main(){
	x := 3
	fmt.Println("x = ", x)

	x1 := add1(&x)
	fmt.Println("x+1 = ", x1)
	fmt.Println("x = ", x) //此时x变了


}

