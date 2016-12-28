// template 模板
package main
import (
    "fmt"
    "html/template"
    "net/http"
)

func main(){
    http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":12345", nil)
    if err != nil {
	fmt.Println(err)
    }
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-type","text/html;charset=utf-8")
    strTemplate := "<div><h3>欢迎光临!</h3></div>"
    t := template.New("test")
    //解析模板
    t,err := t.Parse(strTemplate)
    if err != nil {
	fmt.Println(err)
	return
    }
    err = t.Execute(w, nil)
    if err != nil {
	fmt.Println(err)
    }

}
