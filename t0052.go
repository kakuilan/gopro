//go程序的一般结构
package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

//innitialization of package
func init() {
	println("package init")
}

func main() {
	var a int
	Func1()
	//...
	fmt.Println(a)
}

func Func1() {
	//...
}
