//request.Form是一个url.Values类型,里面存储的是对应的类似key=value的信息
//下面展示了可以对form数据进行的一些操作
package main
import (
		"fmt"
		"net/http"
		)


func main(){
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])


}
