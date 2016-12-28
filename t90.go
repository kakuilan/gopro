package main
import (
    "fmt"
    "time"
)

var wormhole chan time.Time

func deepspace(){
    wormhole <- time.Now()
}

func main(){
    wormhole = make(chan time.Time)
    go deepspace()
    fmt.Println(<-wormhole)
}

