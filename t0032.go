//使用buf实现ioutil.ReadFile
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	content := make([]byte, 0)
	buf := make([]byte, 16)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}

		if n == 16 {
			content = append(content, buf...)
		} else {
			content = append(content, buf[0:n]...)
		}
	}
	fmt.Println(string(content))
}
