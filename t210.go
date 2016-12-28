// 多行字符串
package main
import "fmt"

func main(){
    //错误的
    //s := "Starting part"
    //+ "Ending part"
    
    //正确的
    s := "Starting part" +
	"Ending part"
    //或者
    s2 := `
	Starting part
	Ending part
	`

    fmt.Println(s, s2)

}
