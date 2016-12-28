//调用mymath包中的函数
package main
import (
		"fmt"
		"mymath"
		)

func main(){
	var a,b int = 2,3
	add := mymath.Add(a,b)
	fmt.Println("加法:", add)

	sub := mymath.Sub(a,b)
	fmt.Println("减法:", sub)

	mult := mymath.Mult(a,b)
	fmt.Println("乘法:", mult)

	div := mymath.Div(a,b)
	fmt.Println("除法:", div)



}
