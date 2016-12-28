//https
package main
import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main(){
    url := "https://localhost:10860"
    tr := &http.Transport{
	TLSClientConfig : &tls.Config{
	    InsecureSkipVerify : true,
	},
    }
    c := &http.Client{Transport : tr}
    r, err := c.Get(url)
    if err != nil {
	fmt.Println(err)
	return
    }

    defer r.Body.Close()
    fmt.Println("Header:")
    for k,v := range r.Header {
	fmt.Printf("%s : %s\n", k, v)
    }
    body,err := ioutil.ReadAll(r.Body)
    if err != nil {
	fmt.Println(err)
	return
    }
    fmt.Printf("Body: \n%s\n", body)
}
