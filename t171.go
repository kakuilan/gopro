//遍历目录文件
package main
import (
    "fmt"
    "os"
)

func main(){
    f,err := os.OpenFile(".", os.O_RDONLY, 0666)
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    arrFile, err1 := f.Readdir(0)
    if err1 != nil {
	fmt.Println(err1.Error())
	return
    }

    for k,v := range arrFile {
	fmt.Println(k,"\t", v.Name(), "\t", v.IsDir())
    }
}
