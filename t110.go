//http客户端
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main(){
    r,err := http.Get("http://localhost:1234/图灵")
    if err != nil {
	fmt.Println(err)
    }

    defer r.Body.Close()

    fmt.Println("Header:")
    for k,v := range r.Header {
	fmt.Printf("%s :%s\n", k,v)
    }
    body,err := ioutil.ReadAll(r.Body)
    if err != nil {
	fmt.Println(err)
    }
    fmt.Printf("Body:\n%s\n", body)

}
