//字符串类型
package main

import "fmt"

func main() {
	var b1 byte = 65
	var b2 byte = '\x41'
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'

	fmt.Printf("%d - %d\n", b1, b2)
	fmt.Printf("%c - %c\n", b1, b2)
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3)
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)

}
