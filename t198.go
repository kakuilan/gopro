// http web server
package main
import (
    "net/http"
)

func main(){
    http.HandleFunc("/test", HandleRequest)
    http.ListenAndServe(":8888", nil)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>第一个GO WEB应用</h1>"))
    w.Write([]byte(r.URL.Path))
}
