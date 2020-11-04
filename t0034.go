//创建文件
package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	path := "txt.log"

	newFile, err := os.Create(path)
	checkErr(err)
	defer newFile.Close()

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file doesn`t exist!")
			return
		}
	}

	fmt.Println("file does exist,file name:", fileInfo.Name())
}
