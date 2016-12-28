//http client
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main(){
    host := "http://localhost:1234"
    user := url.Values{
	"Name" : {"Megan Fox"},
	"Sex" : {"female"},
    }

    q := host + "?" + user.Encode()
    fmt.Println(q)
    r,err := http.Get(q)
    if err != nil {
	fmt.Println(err)
	return
    }

    defer r.Body.Close()
    body,err := ioutil.ReadAll(r.Body)
    if err != nil {
	fmt.Println(err)
	return
    }
    fmt.Printf("%s", body)
}
