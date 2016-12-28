//map
package main
import "fmt"

func main(){
	//初始化一个字典
	rating := map[string]float32{"C":5, "GO":4.5, "Python":5, "C++":2}
	
	//map有2个返回值,第二个返回值,如果不存在key,那么ok为false,否则为true
	csharp,ok := rating["C#"]
	if ok{
		fmt.Println("C# is in the map and its rating is ", csharp)
	}else{
		fmt.Println("We have no rating associated with  C# in the map")
	}

	//删除key为C的元素
	delete(rating, "C")


}


