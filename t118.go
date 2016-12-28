//gob 存数据
package main
import (
    "encoding/gob"
    "fmt"
    "os"
)

func main(){
    GDP := map[string] float64 {
	"USA" : 14.58,
	"China" : 5.92,
	"Japan" : 5.45,
    }

    name := "test.gob"
    file,err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
    defer file.Close()
    if err != nil {
	fmt.Println(err)
    }

    enc := gob.NewEncoder(file)
    if err := enc.Encode(GDP);err !=nil {
	fmt.Println("Cannot encode:", err)
	return
    }

}
