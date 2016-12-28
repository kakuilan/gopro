package main
import "fmt"

func main(){
    GDPof := map[string]float64{
	"USA": 14.58,
	"Japan" : 5.45,
    }
    GDPof["China"] = 5.92

    for k,v := range GDPof{
	fmt.Printf("%s:%g\n", k,v)
    }

}
