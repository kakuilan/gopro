package main
import "fmt"

func main(){
    x := 24
    v := x
    if v == nil {
	fmt.Println("x is nil")
    } else if i, is_int := v.(int);is_int {
	fmt.Println("i is an int ", i)
    } else if i, is_float64 := v.(float64); is_float64 {
	fmt.Println("i is a float64 ", i)
    } else if i, is_func := v.(func(int) float64); is_func{
	fmt.Println("i is a function")
    } else {
	i1, is_bool := v.(bool)
	i2, is_string := v.(string)
	if is_bool || is_string {
	    i := v
	    fmt.Println("type is bool or string")
	}else {
	    i:= v
	    fmt.Println("don`t know the type")
	}
    }
}
