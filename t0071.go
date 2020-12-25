//数组的定义和遍历
package main

import "fmt"

func main() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d \n", i, arr1[i])
	}

	println()
	for i, v := range arr1 {
		fmt.Printf("Array at index %d is %d \n", i, v)
	}

}
