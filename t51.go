package main
import (
    "net/http"
    "os"
    "io"
    "fmt"
)

func main(){
    resp,err := http.Get("http://baidu.com")
    if err != nil {
	fmt.Println("Error.")
    }
    defer resp.Body.Close()
    io.Copy(os.Stdout, resp.Body)
}
