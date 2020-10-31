//任意位置读文件f.ReadAt
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	b1 := make([]byte, 2)
	f.ReadAt(b1, 5)
	fmt.Println(string(b1))

}
