//使用append增加切片元素
package main
import "fmt"

func main(){
	//使用make创建切片,len=3,cap=6
	var slice1 = make([]int, 3, 6)
	//给切片增加元素且未超出cap
	slice2 := append(slice1, 1,2,3)
	//给切片增加元素且超出cap
	slice3 := append(slice1, 1,2,3,4)

	fmt.Printf("len=%d cap=%d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d %v\n", len(slice3), cap(slice3), slice3)

}
