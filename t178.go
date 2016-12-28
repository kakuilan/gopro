// xml 结构点绑定处理
package main
import (
    "encoding/xml"
    "fmt"
)

type Student struct {
    XMLName xml.Name `xml:"student"`
    Name string `xml:"name,arrt"`
    Age int `xml:"age,attr"`
    Phone []string `xml:"phones>phone",`
}

type ABC string
func main(){
    str := `
	<?xml version="1.0" encoding="utf-8" ?>
	<student name="张龙" age="27" >
	<phones>
	    <phone>13812345678</phone>
	    <phone>654321</phone>
	</phones>
	</student>`
    var s Student
    xml.Unmarshal([]byte(str), &s)
    fmt.Println(s)


}
