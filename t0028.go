//任意位置读文件f.Seek+f.Read
package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	b1 := make([]byte, 2)

	f.Seek(5, 0)
	f.Read(b1)
	fmt.Println(string(b1))

}
