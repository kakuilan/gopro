package main
import "fmt"

func main(){
    x := 123
    switch i := x.(type) {
	case nil:
	    printString("x is nil")
	case int:
	    printInt(i)
	case float64:
	    printFloat64(i)
	case func(int) float64:
	    printFunction(i)
	case bool, string:
	    printString("type is bool or string")
	default:
	    printString("don`t konw the type")
    }
}
