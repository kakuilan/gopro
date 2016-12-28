//map字典

package main
import "fmt"

func main(){
    mp := make(map[string]string) //key是字符串类型,值也是字符串类型
    mp["a"] = "1"
    mp["b"] = "2"
    mp["pi"] = "3.1415926"
    mp["sh"] = "上海"

    v,ok := mp["sh"] //sh存在,v是value值,ok值为true
    if ok {
	fmt.Println(v)
    } else {
	fmt.Println("key sh 不存在")
    }

    v,ok = mp["bj"] //bj不存在,ok为false
    if ok {
	fmt.Println(v)
    }else{
	fmt.Println("key bj 不存在")
    }


}
