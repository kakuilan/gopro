//接口-打印不通几何图形的面积和周长
package main

import (
	"fmt"
	"math"
)

//内嵌接口
type tmp interface {
	area() float32
}

type geometry interface {
	//	area() float32
	tmp
	perim() float32
}

type rect struct {
	len, wid float32
}

func (r rect) area() float32 {
	return r.len * r.wid
}

func (r rect) perim() float32 {
	return 2 * (r.len + r.wid)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func show(name string, param interface{}) {
	switch param.(type) {
	case geometry:
		//类型断言
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("area of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type!")
	}
}

func main() {
	rec := rect{
		len: 1,
		wid: 2,
	}
	show("rect", rec)

	cir := circle{
		radius: 1,
	}
	show("circle", cir)

	show("test", "test param")

}
