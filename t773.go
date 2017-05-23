//defer
package main
import (
    "os"
    "log"
)

func main() {
    var file *os.File 
    var err error
    filename := "./t177.go"
    if file,err = os.Open(filename);err!=nil {
        log.Println("failed to open the file", err)
    }

    defer file.Close()
}
