//字符串方法

package main
import "fmt"

func main(){
	phrase := "hello world!"
	fmt.Printf("string:\"$s\"\n", phrase)
	fmt.Println("index rune char bytes")
	for index,char := range phrase {
		fmt.Printf("%-2d	%U	'%c'	%X\n", index, char, char, []byte(string(char)))
	}

}

