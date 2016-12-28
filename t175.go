// xml Marshal 将对象编码为xml格式
package main
import (
    "encoding/xml"
    "fmt"
)

type Student struct {
    Name string
    Age int
}

func main(){
    s := &Student{"赵武",24}
    result,err := xml.Marshal(s)
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    fmt.Println(string(result))

}
