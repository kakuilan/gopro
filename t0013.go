//使用switch对interface{}进行断言
package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v is an int\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	default:
		fmt.Printf("%v's type is unkonw\n", v)
	}
}

func main() {
	checkType(7)
	checkType("hello")
	checkType([]int{})
}
