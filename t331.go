// 切片
package main
import "fmt"

func main(){
	//直接创建
	s := make([]int, 10) //长度10
	fmt.Println(s, len(s), cap(s))

	//或
	s2 := []int{1,2,3,4,5} //直接创建5个元素的
	fmt.Println(s2, len(s2), cap(s2))

	//间接创建
	var a [10]int
	s3 := a[1:5] //a[1],a[2],a[3],a[4],不包含a[5]
	fmt.Println(s3, len(s3), cap(s3))


	s4 := a[0:8] //a[0]-a[7]
	fmt.Println(s4, len(s4), cap(s4))


}
