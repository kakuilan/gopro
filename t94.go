//通道变量
package main
import "fmt"

type Read struct {
    key string
    reply chan<- string
}

type Write struct {
    key string
    val string
}

var hole = make(chan interface{})

func deepspace(){
    m := map[string] string{}
    for {
	switch r := (<-hole).(type) {
	case Read:
	    r.reply <- m[r.key] + " from Mars."
	case Write:
	    m[r.key] = r.val
	}
    }

}

func main(){
    go deepspace()

    hole <- Write{"Name", "Martin"}
    home := make(chan string)
    hole <- Read{"Name", home}
    fmt.Println(<-home)
}
