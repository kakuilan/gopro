//多选项case语句
package main
import "fmt"

func main(){
	var searchstr string
	//搜索字符串变量
	fmt.Println("请输入要搜索的关键词:")
	fmt.Scanf("%s", &searchstr)

	switch searchstr {
		case "c","C":
			fmt.Println("C programing language")
		case "go","golang":
			fmt.Println("Go programing language")
		case "java":
			fmt.Println("Java programing language")
		default:
			fmt.Println("Not find the result!")
	
	}

}
