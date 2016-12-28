// map, 引用类型,哈希表.
package main
import "fmt"

func main(){
	m := map[int]struct {
		name string
		age int
	}{
		1 : {"user1", 10}, //可省略元素类型
		2 : {"user2", 20}, 
	}

	fmt.Println(m[1].name)
	

}
