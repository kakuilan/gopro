//多个Method可以同名
package main
import (
		"fmt"
		"math"
		)

type rectangle struct {
	width int
	height int
}

type circle struct {
	radius float32
}

//此处area是对象rectangle的方法
func (recv rectangle) area() int{
	return recv.width * recv.height
}

//此处area是对象circle的方法
func (recv circle) area() float32{
	return recv.radius * recv.radius * math.Pi
}

func main(){
	r1 := rectangle{4,3}
	c1 := circle{5}

	fmt.Println(r1.area(), c1.area())


}
