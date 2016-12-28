// 简单的http get例子
package main
import ("io/ioutil";"net/http";"fmt")

func main(){
    r,err := http.Get("http://www.google.com/robots.txt")
    if err != nil {
	fmt.Println(err)
	return
    }
    b,err := ioutil.ReadAll(r.Body) //将整个内容读入b
    r.Body.Close()
    if err == nil {
	fmt.Printf("%s\n", string(b))
    }


}
