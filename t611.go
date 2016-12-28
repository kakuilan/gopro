//slice作为函数参数
//对形参中slice元素的修改,同时也会修改到底层数组元素的值
package main
import "fmt"

func main(){
	var b = [5]int{1,2,3,4,5}
	var s1 []int = b[:]
	f5(s1)
	fmt.Println(b)
	fmt.Println(s1)
}

func f5(s []int) {
	s[0] +=1
	fmt.Println(s[0])
}
