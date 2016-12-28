//使用条件循环近似求pi圆周率
//基本算法:pi/4约等于1-1/3+1/5-1/7...,知道某一项的绝对值小于10^-6为止
package main
import (
		"fmt"
		"math"
		)

func main(){
	var s int =1
	var n,t,pi float64 = 1.0,1.0,0
	for math.Abs(t) > 1e-6 {
		pi = pi +t
		n = n+2
		s = -s
		t = float64(s) /n
	}
	pi = pi * 4
	fmt.Printf("pi=%10.6f\n", pi)
}

