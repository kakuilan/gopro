// xml UnMarshal 将xml字符串反序列化为对象
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
    str := `
	<?xml version="1.0" encoding="utf-8" ?>
	<Student>
	<Name>展昭</Name>
	<Age>29</Age>
	</Student>`
    var s Student
    xml.Unmarshal([]byte(str), &s)
    fmt.Println(s)

}
