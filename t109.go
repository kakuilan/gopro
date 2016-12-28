//http
package main
import (
    "fmt"
    "net/http"
)

func hi(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "%s 你好", r.URL.Path[1:])
}

func main(){
    http.HandleFunc("/", hi)
    http.ListenAndServe(":1234", nil)
}
