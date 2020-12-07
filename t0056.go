//判断字符串开头和结尾
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? does the string \"%s\" have prefix %s\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("T/F? does the string \"%s\" have suffix %s\n", str, "string")
	fmt.Printf("%t\n", strings.HasSuffix(str, "string"))
}
