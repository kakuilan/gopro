package main
import (
    "bytes"
    "strings"
    "fmt"
)

func main(){
    s := "你GO了吗？"
    fmt.Println(bytes.Count([]byte(s), nil))
    fmt.Println(strings.Count(s, ""))
    fmt.Println(bytes.Split([]byte(s), nil))
    fmt.Println(strings.Split(s, ""))
}
