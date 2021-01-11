//map的排序
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34,
		"bravo": 56, "charlie": 23,
		"delta": 78, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v\n", k, barVal[k])
	}

}
