package main
import (
    "fmt"
    "time"
)

type hole chan time.Time

func deepspace(w hole, h int) {
    defer close(w)
    for ;h>0;h-- {
	w <- time.Now()
	time.Sleep(time.Second)
    }
}

func consumer(w hole) {
    for msg := range w{
	fmt.Println("consumer", msg)
    }
}

func main(){
    w := make(hole)
    go deepspace(w, 8)
    go consumer(w)

    for {
	msg, ok := <-w
	if !ok {
	    break
	}
	fmt.Println("main", msg)
    }
}
