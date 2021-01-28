//结构体方法

package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 8
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B //b1是值
	b1.change()
	println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	println(b2.write())
}
