//字符串截取

package main
import (
	"strings"
	"fmt"
)

func main(){
	line := "this is a string line."
	i := strings.Index(line, " ") //获取第一个空格的索引位置
	firstWord := line[:i] //从第一个字开始时切片直到第一个空格
	j := strings.LastIndex(line, " ") //获取最后一个空格
	lastWord := line[j+1:]
	fmt.Println(firstWord, lastWord)




}
