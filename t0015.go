//panic和recover
package main

import (
	"errors"
	"fmt"
)

func panicFunc() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover panic.")
		}
	}()

	//抛出异常
	panic(errors.New("This is test for panic."))
}

func main() {
	fmt.Println("before panic")
	panicFunc()
	fmt.Println("after panic")
}
