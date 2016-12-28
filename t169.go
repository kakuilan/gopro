// ioutil ReadFile
package main
import (
    "fmt"
    "io/ioutil"
)

func main(){
    buf,err := ioutil.ReadFile("test.txt")
    if err != nil {
	fmt.Println(err.Error())
	return
    }

    fmt.Println(string(buf))

}
