//字典的初始化与创建
package main
import "fmt"

func main(){
	//字典声明后如果没有初始化,值为nil
	var map1 map[string]int
	fmt.Println(map1)

	//未初始化而使用 map1["key1"]=1 语句将编译错误
	//map1["key1"] = 1
	
	//使用make创建后就可以给字典增加数据项了
	map1 = make(map[string]int)
	map1["key1"] = 1
	fmt.Println(map1)

	//使用"{}"初始化后也可以给字典增加数据项
	var map2 = map[string]int {}
	map2["key2"] = 2
	fmt.Println(map2)

}
