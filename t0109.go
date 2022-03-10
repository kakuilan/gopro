//从键盘读取输入
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)

	switch input {
	case "Philip\n":
		fmt.Println("Welcome", input)
	case "Chris\n":
		fmt.Println("Welcome", input)
	case "Ivo\n":
		fmt.Println("Welcome", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}

}
