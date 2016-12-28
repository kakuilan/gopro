// http web
package main
import (
    "net/http"
    "fmt"
)

func main(){
    hadler := &HttpHandler{}
    http.ListenAndServe(":8888", hadler)
}

type HttpHandler struct {
}

func (this *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>在结构体HttpHandler里实现ServeHTTP接口<h1>"))
    w.Write([]byte(r.URL.Path))
    w.Write([]byte(fmt.Sprintf("%v", r.URL.Query())))
}
