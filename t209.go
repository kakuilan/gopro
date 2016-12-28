// 字符串改变
package main

import "fmt"

func main() {
	s := "hello"
	c := []rune(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s, s2)
}
