//面向对象
package main
import "fmt"

type rectangle struct {
	width int
	height int
}

//rectangle的方法
func (recv rectangle) area() int {
	return recv.width * recv.height
}

func main(){
	r1 := rectangle{4,3}
	r2 := rectangle{30,15}

	fmt.Println(r1.area(), r2.area())

}
