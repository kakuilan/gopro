// template语法
package main
import (
    "fmt"
    "html/template"
    "os"
)

func main(){
    strTpl := "你好,{{.}}\r\n"
    t,err := template.New("test").Parse(strTpl)
    if err != nil {
	fmt.Println(err)
	return
    }

    //将要输出的数据传到模板中
    err = t.Execute(os.Stdout, "张三丰")
    if err != nil {
	fmt.Println(err)
    }

    strTpl2 := "姓名:{{.Name}}\r\n年龄:{{.Age}}\r\n"
    user := make(map[string] interface{})
    user["Name"] = "吕洞宾"
    user["Age"] = 3000
    t,err = template.New("test2").Parse(strTpl2)
    if err != nil {
	fmt.Println(err)
	return
    }
    //这里用的是一个map,如果改成struct也可以
    err = t.Execute(os.Stdout, user)
    if err != nil {
	fmt.Println(err)
    }

}
