// reslice,就是基于已有slice创建新slice对象,以便在cap允许范围内调整属性
package main
import "fmt"

func main(){
	s := []int {0,1,2,3,4,5,6,7,8,9}

	s1 := s[2:5]
	s2 := s1[2:6:7]
	//s3 := s2[3:6] //Error: slice bounds out if range

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	//fmt.Println(s3)


}
