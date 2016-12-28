//https
package main
import (
   "log"
    "net/http"
)

func hi(w http.ResponseWriter, _ *http.Request){
    w.Write([]byte("你好秘密"))
}

func main(){
    http.HandleFunc("/", hi)
    log.Println("URL:")
    log.Fatal(http.ListenAndServeTLS(":10860", "cert.pem", "key.pem", nil))
}
