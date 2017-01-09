//strings
package main
import (
    "fmt"
    "strings"
)

func main(){
    asciiOlny := func(char rune) rune {
        if  char > 127 {
            return  '?'
        }
        return  char
    }

    fmt.Println(strings.Map(asciiOlny, "Jerome ·osterreich"))
}
