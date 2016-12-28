// ioutil WriteFile
package main
import (
    "fmt"
    "io/ioutil"
)

func main(){
    err := ioutil.WriteFile("test2.txt", []byte("abcdefg"), 0777)
    if err != nil {
	fmt.Println(err.Error())
    }else{
	fmt.Println("ok")
    }

}
