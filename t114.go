//http server
package main
import (
    "fmt"
    "net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "RawQuery: %s\n", r.URL.RawQuery)
    fmt.Fprintf(w, "Name: %s\n", r.URL.Query()["Name"])
}

func main(){
    http.HandleFunc("/", hi)
    http.ListenAndServe(":1234", nil)
}
