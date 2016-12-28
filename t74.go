package main
import (
    "os"
    "fmt"
)

func main(){
    file, err := os.Open("Nepo")
    if err != nil {
	fmt.Println(err)
    }
    fmt.Println(file)
}
