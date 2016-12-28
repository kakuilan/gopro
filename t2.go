package main
import "fmt"

func GetName()(firstName, lastName, nickName string){
    return "May", "Chan", "Chibi Maruko"
}

func main(){
    var v1 int = 10
    var v2 = 10
    v3 := 10
    v4:=2
    _,_,nickName := GetName()

    fmt.Println(v1, v2, v3, v4, nickName);

}
