//gob 读数据
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var GDP map[string]float64

func main() {
	name := "test.gob"
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		fmt.Println("Cannot open:", err)
		return
	}

	dec := gob.NewDecoder(file)
	if err := dec.Decode(&GDP); err != nil {
		fmt.Println("Cannot decode:", err)
		return
	}

	fmt.Println(GDP)
	fmt.Printf("%x\n", &GDP)

}
