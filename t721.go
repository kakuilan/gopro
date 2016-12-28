//在任意空白符分隔字的情况下找出其第一和最后一个字
package main
import (
	"strings"
	"unicode"
	"unicode/utf8"
	"fmt"
)

func main (){
	line := "lal la asd are\u2028fdin	kubg hah"
	i := strings.IndexFunc(line, unicode.IsSpace)
	firstWord := line[:i]
	j := strings.LastIndexFunc(line, unicode.IsSpace)
	_,size := utf8.DecodeRuneInString(line[j:])
	lastWord := line[j+size:]
	fmt.Println(i,j, firstWord, lastWord)


}


