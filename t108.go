//HTTP服务器
package main
import (
    "io"
    "net/http"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "Hi there!")
    })
    http.ListenAndServe(":1234", nil)
}
