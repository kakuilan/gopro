//切片的读写操作实际目标是底层数组,注意索引号的区别
package main
import "fmt"

func main(){
	data := [...]int {0,1,2,3,4,5,6,7,8,9}
	
	s := data[2:4]
	s[0] += 100
	s[1] += 200 //将改变data的元素

	fmt.Println(data)
	fmt.Println(s)


}

