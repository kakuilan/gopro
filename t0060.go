//通过指针修改变量
package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s

	fmt.Printf("string s origin: %s\n", s)

	//修改变量
	*p = "hello world"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)

}
