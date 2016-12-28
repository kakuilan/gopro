//gob
package main
import (
    "fmt"
    "os"
)

func main(){
    name := "test.gob"
    file,err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0666)
    defer file.Close()

    if err != nil {
	fmt.Println(err)
    }


}
