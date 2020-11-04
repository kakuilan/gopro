//WriteAt在文件指定位置写入内容
package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	newStr := "world"

	newFile, err := os.Create(path)
	checkErr(err)

	n1, err := newFile.WriteString(str)
	checkErr(err)
	fmt.Println("n1:", n1)
	readFile(path)

	n2, err := newFile.WriteAt([]byte(newStr), 6)
	checkErr(err)
	fmt.Println("n2:", n2)
	readFile(path)

	n3, err := newFile.WriteAt([]byte(newStr), 0)
	checkErr(err)
	fmt.Println("n3:", n3)
	readFile(path)

}
