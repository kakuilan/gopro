//使用sync.WaitGroup实现并发同步
package main

import (
	"log"
	"sync"
	"time"
)

func doSomething(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("berfore do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go doSomething(1, &wg)
	go doSomething(2, &wg)
	go doSomething(3, &wg)

	wg.Wait()
	log.Printf("finish all jobs \n")
}
