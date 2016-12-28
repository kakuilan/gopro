//指针变量的定义和引用
package main
import "fmt"

func main(){
	var i int //定义整型变量i
	var i_pointer *int //定义执行整型变量的指针 i_pointer
	i = 100 //将100存放到i的内存单元中
	i_pointer = &i //将i的内存地址存放到指针变量i_pointer中
	
	i2 := * i_pointer //通过指针变量i_pointer值存放的地址读取数据
	fmt.Println(i2)

	* i_pointer = 250
	fmt.Println(i, i2)

	//取整
	i3 := 6/5
	fmt.Println(i3)

}
