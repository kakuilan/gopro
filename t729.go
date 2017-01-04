package main
import "fmt"

func main(){
    s := "Alias<->Synonym"
    chars := []rune(s)
    bytes := []byte(s)

    fmt.Printf("%T: %v\n%T: %v\n", chars, chars, bytes, bytes)

}
