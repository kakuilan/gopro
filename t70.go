package main
import "fmt"

const (
    China int = iota
    USA
    Japan
    Total
)

type GDP struct {
    k string
    v float64
}

func main(){
    GDPof := make([]GDP, Total)
    GDPof[China] = GDP{"China", 5.92}
    GDPof[USA] = GDP{"USA", 14.58}
    GDPof[Japan] = GDP{"Japan", 5.45}

    for k,v := range GDPof {
	fmt.Printf("%d. %s:%g\n", k, v.k, v.v)
    }
}
