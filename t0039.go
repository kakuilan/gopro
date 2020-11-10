//json序列号和反序列化
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var (
		data  = `1`
		value int
	)

	err1 := json.Unmarshal([]byte(data), &value)

	fmt.Println("Unmarshal error is:", err1)
	fmt.Printf("Unmarshal value is:%T, %d\n", value, value)

	value2, err2 := json.Marshal(value)
	fmt.Println("Marshal error is:", err2)
	fmt.Printf("Marshal value is:%s\n", string(value2))

}
