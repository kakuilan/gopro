//多重继承
package main

import (
	"fmt"
)

type Camera struct{}

func (c *Camera) TakePicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakePicture())
	fmt.Println("It works like a Phone too: ", cp.Call())

}
