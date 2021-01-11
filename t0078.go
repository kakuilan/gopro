//map类型的切片
package main

import "fmt"

func main() {
	//版本A
	items := make([]map[int]int, 5)
	//通过索引,使用切片的map元素
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)


	//版本B,不好的
	items2 := make([]map[int]int, 5)
	//这里的item项只是map值的一个拷贝
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)

}
