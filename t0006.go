package main

import "fmt"
import "sync"

func main() {
	cMap := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		cMap["北京"] = 1
		wg.Done()
	}()

	go func() {
		cMap["上海"] = 2
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("%v", cMap)
}
