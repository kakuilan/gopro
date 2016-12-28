package main
import "fmt"

//PersionInfo是一个包含个人信息的类型
type PersonInfo struct{
    ID string
    Name string
    Address string
}

func main(){
var myMap map[string] PersonInfo
myMap = make(map[string] PersonInfo, 100)
myMap2 := map[string] PersonInfo{
    "1" : PersonInfo{"1", "Jack", "Room 101,..."}, //必须有,号，否则提示syntax error: need trailing comma before newline in composite literal
}
myMap2["2"] = PersonInfo{"2", "Tom", "Room 203,.."}
myMap2["3"] = PersonInfo{"3", "Hum", "Room 102,.."}
fmt.Println(myMap2)

//删除map内的元素
delete(myMap2, "2")
fmt.Println(myMap2)


fmt.Println(myMap, myMap2)
}
