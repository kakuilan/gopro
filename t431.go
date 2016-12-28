// slice切片
package main
import "fmt"

func main(){
	slice := []byte{'a', 'b', 'c', 'd'}
	var ar = [10]byte{'a','b','c','d','e','f','g','h','i','j'}

	//声明2个含有byte的slice
	var a,b []byte

	//a指向数组的第3个元素开始,并到第5个元素结束
	a = ar[2:5]

	//b是数组ar的另一个slice
	b = ar[3:5]

	fmt.Println(slice)
	fmt.Println(ar)
	fmt.Println(a,b)


}
