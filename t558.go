//bytes.Repeat
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	b := []byte("na")
	count := 2
	fmt.Println("ba" + string(bytes.Repeat(b, count)))
}
