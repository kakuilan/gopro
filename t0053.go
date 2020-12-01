//init函数例子
package main

import (
	"gopro/mypkg"
	"fmt"
)

var twoPi = 2 * mypkg.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)
}
