//channel通道
package main

func main() {
	ch := make(chan string)
	go func() {
		ch <- "ping"
	}()

	println(<-ch)

}
