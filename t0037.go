//通过Buffered Writer写文件
package main

import (
	"bufio"
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

	newFile, err := os.Create(path)
	checkErr(err)
	defer newFile.Close()

	bufferWriter := bufio.NewWriter(newFile)
	for _, v := range str {
		written, err := bufferWriter.WriteString(string(v))
		checkErr(err)
		fmt.Println("Written:", written)
	}

	//此时还读不到,因为还在缓冲里
	readFile(path)

	//检查有多少在缓冲里
	unflushSize := bufferWriter.Buffered()
	fmt.Println("unflushSize:", unflushSize)

	//把缓冲写到磁盘
	bufferWriter.Flush()

	//再读就有了
	readFile(path)
}
