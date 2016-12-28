// range 迭代器
package main
import "fmt"

func main(){
    list := []string {"a", "b", "c", "d", "e", "f"}
    for k,v := range list {
	fmt.Println(k,v)
    }

    //字符串的range
    str :="Hello World.你好，世界"
    for pos,char := range str {
	//fmt.Println(pos, char)
	fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
    }

}
