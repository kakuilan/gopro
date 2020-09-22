//结构体内嵌本身
package main

import "fmt"

type Tree struct {
	value       int
	left, right *Tree
}

func main() {
	tree1 := Tree{
		value: 2,
	}

	tree2 := Tree{
		value: 1,
	}

	fmt.Printf(">>> %#v\n", tree1 == tree2)
}
