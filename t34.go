package main
import "fmt"

func say(str string, args ... interface{})(int,error){
    _,err := fmt.Printf(str, args...)
    return len(args), err
}

func main(){
    count := 1
    //匿名函数
    closure := func(msg string){
	say("%d %s\n", count, msg)
        count++
    }

    closure("Say one");
    closure("Say again");
}

