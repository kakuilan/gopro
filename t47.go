package main
import "fmt"
import "time"

func main(){
    timeout := make(chan bool, 1)
    go func(){
	time.Sleep(1e9)
    }()

    select {
	case <-ch:
	case <-timeout:
	
    }

}
