//http post
package main
import (
    "log"
    "net/http"
    "text/template"
    "time"
)

type User struct {
    Name string
    Date time.Time
}

var users []User

var homePage = template.Must(template.New("home").Parse(
`<html><body>
{{range .}}
{{.Name}} dates on {{.Date.Format "3:04pm, 2 Jan"}} <br />
{{end}}

<form action="/post" method="post" >
姓名:<input type='text' name='Name' /><br />
<input type='submit' value='提交' />
</form></body></html>
`))

func home(w http.ResponseWriter, _ *http.Request) {
    if err := homePage.Execute(w, users); err != nil {
	log.Printf("%v", err)
    }
}

func post (w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm();err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("%v", err)
	return
    }

    users = append(users, User{
	Name : r.FormValue("Name"),
	Date : time.Now(),
    })

    http.Redirect(w, r, "/", http.StatusFound)
}

func main(){
    http.HandleFunc("/", home)
    http.HandleFunc("/post", post)
    log.Println("http://localhost:1234")
    log.Fatalln(http.ListenAndServe(":1234", nil))
}
