package main
import "fmt"

func main(){
    var str string
    str = "Hello world"
    ch := str[0]

    fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
    fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
}
