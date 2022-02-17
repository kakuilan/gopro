//从缓冲读取数据
package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	//从标准输入读取
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}

}
