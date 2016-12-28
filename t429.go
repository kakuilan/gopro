// iota枚举,默认开始值是0,每调用一次加1
package main
import "fmt"

func main(){
	const(
			x = iota
			y = iota
			z = iota
			w //常量声明省略值时,默认和之前一个值的字面相同.这里隐式的说w=iota,因此w==3
			)


	const v = iota //每遇到一个const关键字,iota就会重置,因此v==0
	fmt.Println(x,y,z,w,v)

}
