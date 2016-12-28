//使用for语句统计单词中的某个字符数
package main
import "fmt"

func main(){
	var words string
	var ch byte
	var ln,count int
	fmt.Println("请输入一个英文单词：")
	fmt.Scanf("%s", &words)
	fmt.Println("请输入要统计的英文字符：")
	fmt.Scanf("%c", &ch)
	//计算单词的字节长度
	ln = len(words)

	for i:=0;i<ln;i++{
		if byte(words[i])==ch {
			count++
		}
	}
	fmt.Printf("There are %d %q in the %q\n", count,ch,words)

}
