//切片元素的遍历
package main
import "fmt"

func main(){
	var slice1 = []int{1,2,3,4,5}
	//使用下标访问切片元素
	for i:=0;i<=4;i++{
		fmt.Printf("slice1[%d]=%d ", i, slice1[i])
	}
	fmt.Println()
	//使用range
	for i,v :=range slice1 {
		fmt.Printf("slice1[%d]=%d  ", i,v)
	}
	fmt.Println()

}
