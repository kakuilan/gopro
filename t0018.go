//单向channel
package main

func receive(receiver chan<- string, msg string) {
	receiver <- msg
}

func send(sender <-chan string, receiver chan<- string) {
	msg := <-sender
	receiver <- msg
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	receive(ch1, "pass message")
	send(ch1, ch2)

	println(<-ch2)
}
