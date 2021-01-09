//map操作
package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["NewDelhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25

	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing \n")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("value is: %d\n", value)

	//delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Washington \n")
	}

}
