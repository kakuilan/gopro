//xml序列化和反序列化
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Student struct {
	Name    string `xml:"name"`              //xml标签
	Address string `xml:"address,omitempty"` //若该字段为空就过滤掉
	Hobby   string `xml:"-"`                 //进行xml序列化的时候忽略该字段
	Father  string `xml:"parent>father"`     //xml标签嵌套模式
	Mother  string `xml:"parent>mother"`     //xml标签嵌套模式
	Note    string `xml:"note,attr"`         //xml标签属性
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	stud1 := Student{
		Name:  "haha",
		Hobby: "basketball",
	}

	data, _ := xml.Marshal(stud1)
	fmt.Println(string(data))

	//xml序列化
	newData, err := xml.MarshalIndent(stud1, " ", "  ")
	checkErr(err)
	fmt.Println(string(newData))

	//写xml文件
	err = ioutil.WriteFile("stu.xml", newData, 0644)
	checkErr(err)

	//读xml文件
	content, err := ioutil.ReadFile("stu.xml")
	stu2 := &Student{}

	//xml反序列化
	err = xml.Unmarshal(content, stu2) //注意：第二个参数必须是指针
	checkErr(err)

	fmt.Printf("stu2: %#v\n", stu2)

}
