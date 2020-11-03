//f.ReadAt是并发安全的
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, _ := os.Open("t0001.go")
	defer f.Close()

	for i := 0; i < 5; i++ {
		go func() {
			b1 := make([]byte, 2)

			f.ReadAt(b1, 5)
			fmt.Println(string(b1))

			f.ReadAt(b1, 2)
			fmt.Println(string(b1))
		}()
	}

	time.Sleep(time.Second)
}
