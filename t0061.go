//switch语句
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int = rand.Int()
	switch {
	case num == 98:
		fmt.Println("It`s equal to 98")
	case num == 100:
		fmt.Println("It`s equal to 100")
	case num == 999:
		fallthrough
	case num > 100:
		fmt.Println("It`s greater than 100")
	default:
		fmt.Println("It`s :", num)
	}

}
