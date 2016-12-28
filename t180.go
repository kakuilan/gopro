// json转对象
package main
import (
    "encoding/json"
    "fmt"
)

type Student struct {
    Name string `json:"userName"`
    Age int
}

func main(){
    a := Student{"包公",56}
    s := &Student{"包公",56}
    
    //将s编码为json
    buf,err := json.Marshal(s)
    fmt.Println(a,s, buf,err)
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    fmt.Println(string(buf))
    //将json字符串转换为Student对象
    var s1 Student
    json.Unmarshal(buf, &s1)
    fmt.Println(s1)

}
