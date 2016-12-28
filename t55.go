package main
import "fmt"

func main(){
    var s = "Go程序"
    var r = []rune(s)
    fmt.Printf("%c %c", s[0],s[2])
    fmt.Println("")
    fmt.Printf("%x", s)
    fmt.Println()

    fmt.Printf("%c %c", r[0], r[2])
    fmt.Println()
    fmt.Printf("%x", r)
}
