//接口例子
package main

import "fmt"

//定义接口类型
type Shaper interface {
	Area() float32
}

//结构体
type Square struct {
	side float32
}

//实现方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1

	//或者
	areaIntf2 := Shaper(sq1)

	fmt.Printf("The square has area: %f\n", areaIntf.Area())
	fmt.Printf("The square has area: %f\n", areaIntf2.Area())

}
