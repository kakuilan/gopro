//计算函数执行时间
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	_ = dosomething()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("dosomethin took this amout of time: %s\n", delta)
}

func dosomething() int {
	all := 100000000
	num := 0
	for i := 0; i <= all; i++ {
		num += i + 2
	}
	return num
}
