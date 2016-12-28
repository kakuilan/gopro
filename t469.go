// http session
// 先安装包, go get github.com/gorilla/sessions
package main
import (
		"fmt"
		"net/http"
		"github.com/gorilla/sessions"
		)

var store = sessions.NewCookieStore([]byte("something-very-secret"))


func MyHandler(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r, "session-name")
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Save(r, w)
}


func main(){
	http.HandleFunc("/", MyHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}


}



