//格式化数字
package main
import (
    "fmt"
    "strings"
    "unicode/utf8"
)

func Pad(number, width int, pad rune) string {
    s := fmt.Sprint(number)
    gap := width - utf8.RuneCountInString(s)
    if gap >0 {
        return strings.Repeat(string(pad), gap) + s
    }
    return s
}

func main(){
    //二进制
    fmt.Printf("|%b|%9b|%-9b|%09b|% 9b|\n", 37, 37, 37, 37, 37)
    
    //八进制
    fmt.Printf("|%o|%#o|%# 8o|%#+ 8o|%+08o|\n", 41,41,41,41,-41)

    //十六进制
    i := 3931
    fmt.Printf("|%x|%X|%8x|%08x|%#04X|0x%04X|\n", i,i,i,i,i,i)

    //十进制
    i = 569
    fmt.Printf("|$%d|$%06d|$%+06d|$%s|\n", i,i,i,Pad(i,6, '*'))

    //
    fmt.Printf("%d %#04x %U '%c'\n", 0x3A56, 934, '\u03A6', '\U000003A6')

}

