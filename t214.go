// switch
package main
import "fmt"

func unhex(c byte) byte {
    switch {
	case c==22:
	    fallthrough
	case '0'<=c && c <='9':
	    return c - '0'
	case 'a' <=c && c<='f':
	    return c - 'a' + 10
	case 'A' <=c && c<='F':
	    return c - 'A' +10
    }
    return 0
}

func main(){
    s := byte('$')
    t := unhex(s)
    fmt.Println(s, t)
}
