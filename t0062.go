//Unicode编码的字符串遍历
package main

import "fmt"

func main() {
	//ASCII编码
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("character on position %d is: %c \n", ix, str[ix])
	}

	str2 := "This is 中文字符串！"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("character on position %d is: %c \n", ix, str2[ix])
	}

	//使用for-range
	for pos, char := range str {
		fmt.Printf("-- character on position %d is: %c \n", pos, char)
	}
	println("--------str2:")
	for pos, char := range str2 {
		fmt.Printf("-- character %c starts at byte postion %d \n", char, pos)
	}
	println()
	for index, rune := range str2 {
		fmt.Printf("%-2d	%d	%U '%c'	%X\n", index, rune, rune, rune, []byte(string(rune)))
	}

}
