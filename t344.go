// copy,在两个slice间复制数据,复制长度以len小的为准.两个slice可指向同一底层数组,允许元素区间重叠.
package main
import "fmt"

func main(){
	data := [...]int {0,1,2,3,4,5,6,7,8,9}
	s := data[8:]
	s2 := data[:5]
	fmt.Println(len(s), len(s2))
	fmt.Println(s, s2)

	copy(s2, s) // dst:s2, src:s
	fmt.Println(s2)
	fmt.Println(data)



}
