// 空结构"节省"内存,比如用来实现set数据结构,或者实现没有"状态"只有方法的"静态类"
package main
import "fmt"

var null struct{}

func main(){
	set := make(map[string]struct{})
	set["a"] = null
	
	fmt.Println(set)

}
