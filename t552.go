//字典项的查找
package main
import "fmt"

func main(){
	var map1 = map[string]int{"key1":100, "key2":200}

	//key值存在否
	v,ok := map1["key1"]
	if ok{
		fmt.Println("key1存在", v, ok)
	}else{
		fmt.Println("key1不存在", v,ok)
	}

	//key值不存在
	v,ok = map1["key3"]

	if ok{
		fmt.Println("key3存在", v, ok)
	}else{
		fmt.Println("key3不存在", v,ok)
	}
}
