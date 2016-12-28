//func
package main
import "fmt"

func A(){
    B()
}

func B(){
    fmt.Println("OK")
}

func main(){
    A()
}
