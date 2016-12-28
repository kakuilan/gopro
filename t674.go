//reflect TypeOf
package main
import (
		"fmt"
		"reflect"
		)

func main(){
	var pi float64 = 3.14
	t := reflect.TypeOf(pi)
	v := reflect.ValueOf(pi)

	fmt.Println("Type:", t.Kind())
	fmt.Println("Value:", v.Float())


}
