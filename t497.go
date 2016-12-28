//自定义类型转换
package main
import "fmt"

type(
		newint int
		newfloat float32
		)

func main(){
	var i newint = 100
	var f newfloat
	f = newfloat(i)
	fmt.Printf("%d , %f\n", i, f)
}
