package main
import "fmt"
import "net/url"

func main(){
    var b []byte
    b = append(b, "hi"...)//b == []byte{'h','i'}
    fmt.Println(b)

    gurl, er := url.Parse("http://golang.org/pkg")
    fmt.Println(gurl.Host, er)

}
