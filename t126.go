//if中包含初始化表达式,只能有一个
package main
import "fmt"

func main(){
    if a:=2;a ==4 {
	fmt.Println("OK")
    }else if(a<2){
	fmt.Println("a<2")
    }else{
	fmt.Println("a=", a)
    }

}
