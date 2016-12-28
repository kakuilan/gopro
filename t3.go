package main
import "fmt"

func main(){
    const Pi float64 = 3.141592653
    const zero = 0.0
    const (
	size int64 = 1024
	eof = -1
    )
    const u,v float32 = 0,3
    const a,b,c = 3,4,"foot"

    fmt.Println(Pi, zero, size, eof, u,v, a,b,c)

    const mask = 1<<3
    fmt.Println(mask);

}
