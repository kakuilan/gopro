//json序列号和反序列化
package main

import (
	"encoding/json"
	"fmt"
)

func printHelper(name string, value interface{}) {
	fmt.Printf("%s Unmarshal value is: %T, %v \n", name, value, value)
}

func main() {
	var (
		d1 = `false`
		v1 bool
	)

	json.Unmarshal([]byte(d1), &v1)
	printHelper("d1", v1)

	var (
		d2 = `2`
		v2 int
	)

	json.Unmarshal([]byte(d2), &v2)
	printHelper("d2", v2)

	var (
		d3 = `3.14`
		v3 float32
	)

	json.Unmarshal([]byte(d3), &v3)
	printHelper("d3", v3)

	var (
		d4 = `[1,2]`
		v4 []int
	)

	json.Unmarshal([]byte(d4), &v4)
	printHelper("d4", v4)

	var (
		d5 = `{"a":"b"}`
		v5 map[string]string
		v6 interface{}
	)

	json.Unmarshal([]byte(d5), &v5)
	printHelper("d5", v5)

	json.Unmarshal([]byte(d5), &v6)
	printHelper("d5(interface{})", v6)

}
