// method
package main
import "fmt"
import "math"

type Rectangle struct {
	width,height float64
}

type Circle struct {
	radius float64
}

//不是方法
func area(r Rectangle) float64 {
	return r.width * r.height
}

//这个才是方法
func (r Rectangle) areaMed() float64 {
	return r.width * r.height
}

func (c Circle) areaMed() float64 {
	return c.radius * c.radius * math.Pi
}

func main(){
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9,4}
	c1 := Circle{10}
		
	fmt.Println("用函数求r1的面积是:", area(r1))
	fmt.Println("用方法求r2的面积是:", r2.areaMed())
	fmt.Println("用方法求c1的面积是:", c1.areaMed())
	

}
