//interface接口例子
package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

//定义结构体的方法
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

//定义接口
type valuable interface {
	getValue() float32
}

//参数类型时valuable接口
func showValue(asset valuable) {
	fmt.Printf("value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOD", 577.20, 4}
	showValue(o)

	o = car{"BMW", "M3", 665000}
	showValue(o)
}
