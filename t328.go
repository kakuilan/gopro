// 数组
package main
import "fmt"

func main(){
	a := [3]int{1,2} //未初始化元素的值为0
	b := [...]int{1,2,3,4} //通过初始化值确定数组长度
	c := [5]int{2:100, 4:200} //使用索引初始化元素

	d := [...]struct {
		name string
		age uint8
	}{
		{"user1", 10}, //可省略元素类型
		{"user2", 20}, //但是别忘了最后一行的逗号
	}

	//支持多维数组
	a2 := [2][3]int{{1,2,3}, {4,5,6}}
	b2 := [...][2]int{{1,2}, {2,2}, {3,3}} //第二维度不能用"..."

	fmt.Println(a,b,c,d)
	fmt.Println(a2, b2)


}
