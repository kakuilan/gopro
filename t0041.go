//json序列化
package main

import (
	"encoding/json"
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
	s := Student{
		Id:   1,
		Name: "小红",
		Results: []Result{
			Result{
				Name:  "语文",
				Score: 90,
			},
			Result{
				Name:  "数学",
				Score: 100,
			},
		},
	}

	dat, _ := json.Marshal(s)
	ioutil.WriteFile("data.json", dat, 0755)

}
