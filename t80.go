package main
import "fmt"

//类型断言和判断

type (
    T0 []string
    T00 []string
    T1 struct{a, b int}
    T11 struct{c, d int}
    T2 func(int, float64) *T0
    T22 func(int, float64) *[]string
)

var (
    t interface{}
    t0 T0
    t1 T1
    t2 T2
)

func main(){
    t = t0
    {
	v, ok := t.(T0)
	fmt.Println(v, ok)
    }
    {
	v, ok := t.([]string)
	fmt.Println(v, ok)
    }

    t = t1
    {
	v, ok := t.(T1)
	fmt.Println(v, ok)
    }
    {
	v, ok := t.(T11)
	fmt.Println(v, ok)
    }

    t = t2
    {
	v, ok := t.(T2)
	fmt.Println(v, ok)
    }
    {
	v, ok := t.(T22)
	fmt.Println(v, ok)
    }
}
