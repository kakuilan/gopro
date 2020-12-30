//切片
package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] //不包含索引5的元素

	//给数组元素赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	println()
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//扩充切片
	println()
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice new at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length2 of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity2 of slice1 is %d\n", cap(slice1))

}
