//数组的传值和引用
package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	var ar1 [3]int
	var ar2 = new([3]int) //返回的是指针

	f(ar1)
	fp(&ar1)

	println()
	f(*ar2)
	fp(ar2)

}
