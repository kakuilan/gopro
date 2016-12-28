// 分组声明
package main
import (
		"fmt"
		"errors"
		)

const(
		i = 100
		pi = 3.1415
		prefix = "go_"
		)

var (
		j int
		k float32
		pr string
		)

func main(){
	var e = errors.New("asdfasd")
	fmt.Println(i,pi,prefix)
	fmt.Println(j,k,pr, e)

}
