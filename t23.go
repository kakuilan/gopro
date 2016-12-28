package main
import "fmt"

//PersionInfo是一个包含个人信息的类型
type PersionInfo struct{
    ID string
    Name string
    Address string
}

func main(){
    var persionDB map[string] PersionInfo
	persionDB = make(map[string] PersionInfo)

	//往这个map里插入几条数据
	persionDB["12345"] = PersionInfo{"12345", "Tom", "Rom 203,..."}
	persionDB["1"] = PersionInfo{"1", "Jack", "Rom 101,..."}

	//从这个map查找键为'1234'的信息
	person, ok := persionDB["1234"]
	if ok{
	    fmt.Println("Found person", person.Name, " with ID 1234.")
	}else{
	    fmt.Println("Did not find person with ID 1234")
	}

}
