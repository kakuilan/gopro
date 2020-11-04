//写文件
package main

import (
	"fmt"
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(path string) {
	data, err := ioutil.ReadFile(path)
	checkErr(err)
	fmt.Println("file content:", string(data))
}

func main() {

	path := "test.log"
	str := "hello"
	err := ioutil.WriteFile(path, []byte(str), 0644)
	checkErr(err)
	readFile(path)
}
