//cookie
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

var users = make(map[string] []User)
var homePage = template.Must(template.New("home").Parse(
`
<html><body>
{{range .}}
{{.Name}} dates on {{.Date.Format "3:04pm, 2 Jan"}}<br />
{{end}}

<form action="/post" method="post" >
姓名：<input type='text' name='Name' /><br />
<input type='submit' value='提交' />
</form></body></html>
`)) //注意这里的`和)不能换行写

func home(w http.ResponseWriter, r *http.Request){
    var us []User
    cookie,err := r.Cookie("user")
    if err == http.ErrNoCookie {
	v := time.Now().Format(time.RFC3339)
	http.SetCookie(w, &http.Cookie{
	    Name : "User",
	    Value : v,
	})
	log.Printf("New user: %s", v)
    }else{
	us = users[cookie.Value]
	log.Printf("Return user: %s", cookie.Value)
    }

    if err := homePage.Execute(w, us);err !=nil {
	log.Printf("%v", err)
    }
}

func post (w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("user")
    if err == http.ErrNoCookie {
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("%v", err)
	return
    }

    if err := r.ParseForm();err != nil {
	w.WriteHeader(http.StatusInternalServerError)
	log.Printf("%v", err)
	return
    }

    users[cookie.Value] = append(users[cookie.Value], User{
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
