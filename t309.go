// 函数变参,本质上就是slice,只能有一个,且必须是最后一个
package main
import "fmt"

func test(s string, n ...int) string {
	var x int
	for _,i := range n {
		x += i
	}
	
	return fmt.Sprintf(s, x)
}

func main(){
	println(test("sum: %d", 1,2,3,3))

	//使用slice对象做变参时,必须展开
	s := []int{12,34,6,73}
	println(test("sum: %d", s...))

}

