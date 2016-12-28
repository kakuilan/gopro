// 数组值拷贝会造成性能问题,通常建议使用slice，或数组指针
package main
import "fmt"

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}

func main(){
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test(a)
	fmt.Println(a)

	//内置函数len,cap都返回数组长度
	b := [3]int{}
	println(len(b), cap(b))
}
