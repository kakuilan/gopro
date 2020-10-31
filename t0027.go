//带缓冲区读文件
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	buf := make([]byte, 16)
	f.Read(buf)
	fmt.Println(string(buf))
}
