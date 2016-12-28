//单词和字母统计:1.计算字符数量(包括空格);2.计算单词数量;3.计算行数 换言之,实现一个wc(l)
package main
import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func main(){
    var chars, words, lines int
    r := bufio.NewReader(os.Stdin)
    for {
	switch s,ok := r.ReadString('\n');true {
	case ok != nil :
	    fmt.Printf("%d %d %d\n", chars, words, lines);
	    return
	default:
	    chars += len(s)
	    words += len(strings.Fields(s))
	    lines++
	}

    }



}




