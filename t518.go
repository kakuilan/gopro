//从键盘输入三角形3边长,求三角形面积和周长,输出结果,保留2位有效位
package main
import (
		"fmt"
		"math"
		)


func main(){
	var a,b,c,l,hl,area float64
	fmt.Println("请输入三角形三边：")
	fmt.Scanf("%f,%f,%f", &a,&b,&c)
	l = a+b+c
	fmt.Printf("三角形周长=%6.2f\n", l)
	hl = l * 0.5
	area = math.Sqrt(hl*(hl-a)*(hl-b)*(hl-c)) //开方,平方根
	fmt.Printf("三角形面积=%6.2f\n", area)



}
