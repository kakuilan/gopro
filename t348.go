// 可在map迭代时安全删除键值.但如果期间有新增操作,那么就不知道会有什么意外了.
package main
import "fmt"

func main(){
	for i:=0;i<5;i++{
		m := map[int]string {
			0:"a", 1:"a", 2:"a", 3:"a", 4:"a",
			5:"a", 6:"a", 7:"a", 8:"a", 9:"a",
		}

		for k:= range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	
	}


}
