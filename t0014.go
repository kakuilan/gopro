//defer延迟函数
package main

import "fmt"

//注意比较这2个函数的差异
func printNum1() {
	//defer延迟时,i=5
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func printNum2() {
	//传参
	for i := 0; i < 5; i++ {
		defer func(v int) {
			fmt.Println(v)
		}(i)
	}
}

func main() {
	printNum1()
	println()
	printNum2()
}
