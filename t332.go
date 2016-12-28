// 切片
package main
import "fmt"

func main(){
	data := [...]int{0,1,2,3,4,5,6,7,8,9}
	
	s1 := data[:6:8]
	fmt.Println(s1, len(s1), cap(s1))


	s2 := data[5:]
	fmt.Println(s2, len(s2), cap(s2))


	s3 := data[:3]
	fmt.Println(s3, len(s3), cap(s3))


	s4 := data[:]
	fmt.Println(s4, len(s4), cap(s4))
}
