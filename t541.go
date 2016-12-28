//数组的初始化
package main
import "fmt"

func main(){
	var a [10]int
	var b = [10]int{1,2,3,4,5,6,7,8,9,10}
	var c = [10]int{1,2,3,4,5}
	var d = [10]int{2:4,5:7}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
