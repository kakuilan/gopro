//全量读文件
package main

import (
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("t0001.go")
	if err != nil {
		println(err.Error())
	} else {
		println(string(data))
	}

}
