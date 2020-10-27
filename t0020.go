//select语句
package main

import "time"

func strWorker(ch chan string) {
	time.Sleep(1 * time.Second)
	println("do something with strWorker...")
	ch <- "str"
}

func intWorker(ch chan int) {
	time.Sleep(1 * time.Second)
	println("do something with intWorker...")
	ch <- 1
}

func main() {
	chStr := make(chan string)
	chInt := make(chan int)

	go strWorker(chStr)
	go intWorker(chInt)

	for i := 0; i < 2; i++ {
		select {
		case <-chStr:
			println("get value from strWorker")
		case <-chInt:
			println("get value from intWorker")
		}
	}
}
