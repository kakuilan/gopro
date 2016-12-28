// 可通过meta设置为代码库设置自定义路径
package main
import (
		"fmt"
		"net/http"
		)


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `<meta name="go-import"
			 content="192.168.128.129/qyuhen/test git https://github.com/qyuhen/test"`)
}

func main(){
	http.HandleFunc("/qyuhen/test", handler)
	http.ListenAndServe(":80", nil)


}
