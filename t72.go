package main
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

type Lockmap struct {
    sync.Mutex
    m map[string]int
}
var lockmap = make([]Lockmap, 10)

func init(){
    for i := range lockmap{
    	lockmap[i].m = make(map[string]int)
    }
}

func counter(s string){
    i := int(s[0] - '0')
    time.Sleep(time.Duration(i) * time.Millisecond)

    lockmap[i].Lock()
    defer lockmap[i].Unlock()
    lockmap[i].m[s]++
}

func main(){
    for i := 0; i<20; i++ {
	r := fmt.Sprintf("%d", rand.Uint32())
	go counter(r)
    }

    time.Sleep(time.Second)
    for i := range lockmap{
	for k,v := range lockmap[i].m {
	    fmt.Printf("[%d] %s = %d\n", i, k, v)
	}
    }
}
