//使用bufio实现行统计
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	br := bufio.NewReader(f)
	totalLine := 0

	for {
		_, isPrefix, err := br.ReadLine()
		if !isPrefix {
			totalLine += 1
		}
		if err == io.EOF {
			break
		}
	}

	fmt.Println("total lines is:", totalLine)
}
