//json反序列化
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Result struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type Student struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Results []Result `json:"results"`
}

func main() {
	dat, _ := ioutil.ReadFile("data.json")
	var s Student
	json.Unmarshal(dat, &s)
	fmt.Printf("Student`s result is: %v\n", s)
}
