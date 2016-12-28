// 可直接创建slice对象,自动分配底层数组
package main
import "fmt"

func main(){
	s1 := []int {0,1,2,3,8:100} //通过初始化表达式构造,可使用索引号
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) //使用make创建,指定len和cap
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) //省略cap,相当于cap=len
	fmt.Println(s3, len(s3), cap(s3))



}
