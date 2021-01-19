//结构体
package main

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"

	//初始化结构体
	ms1 := &struct1{8, 3.14, "Zhang3"}
	//或者
	var ms2 struct1
	ms2 = struct1{99, 2.899, "Li4"}

	fmt.Printf("struct1: %v\n", ms)
	fmt.Println(ms1)
	fmt.Println(ms2)

}
