package main
import (
    "fmt"
    "math/rand"
    "time"
)

type hole chan int

func deepspace(w hole, h int){
    defer close(w)
    d := time.Duration(rand.Intn(h)) * time.Second

    for ;h >0;h-- {
	w <- rand.Intn(4)
	time.Sleep(d)
    }
}

func main(){
    n :=8
    w := make(hole)
    t := 0
    maxTime := time.Second

    go deepspace(w, n)
Out:
    for i := 0;i<n;i++{
	select {
	case n := <-w:
	    t += n
	case <-time.After(maxTime):
	    fmt.Println("Time out")
	    break Out
	}
    }
    fmt.Println("Total: ", t)
}
