// append 一旦超出slice.cap限制,就会重新分配底层数组,即便原数组并未填满
package main
import "fmt"

func main(){
	data := [...]int {0,1,2,3,4,5,6,7,8,9}
	s := data[:2:3]

	s = append(s, 100, 200) //一次append两个值,超出s.cap限制
	//重新分配底层数组,与原数组无关
	fmt.Println(s)
	fmt.Println(data)
	fmt.Println(&s[0], &data[0])


}


