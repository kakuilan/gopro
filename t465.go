//简单路由
package main
import (
		"fmt"
		"net/http"
		)


type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w,r)
		return
	}
	http.NotFound(w,r)
	fmt.Fprintf(w, "Router not found!")
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myroute test!")
}

func main(){
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)

}
