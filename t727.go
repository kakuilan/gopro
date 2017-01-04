//fmt 字符串
package main
import (
    "fmt"
)

func main(){
    slogan := "End $& r~ ttl  ";
    fmt.Printf("%s\n%q\n%+q\n%#q\n", slogan, slogan, slogan, slogan)

    chars := []rune(slogan)
    fmt.Printf("%x\n%#x\n%#X\n", chars, chars, chars)

    bytes := []byte(slogan)
    fmt.Printf("%s\n%x\n%X\n% X\n%v\n", bytes, bytes, bytes, bytes, bytes)
}
