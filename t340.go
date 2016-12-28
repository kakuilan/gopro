// append,向slice尾部添加数据,返回新的slice对象
package main
import "fmt"

func main(){
	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)

	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s, s2)



}

