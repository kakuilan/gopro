// http web
// http://192.168.128.129:9090/?url_long=1234&a=qwer&ttl=123
package main
import (
		"fmt"
		"net/http"
		"strings"
		"log"
		)


func sayhelloName(w http.ResponseWriter, r *http.Request ) {
	r.ParseForm() //解析参数,默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello lkk.") //这个写入到w的是输出到客户端
}

func main(){
	http.HandleFunc("/", sayhelloName) //设置访问路由
	err := http.ListenAndServe(":9090", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
