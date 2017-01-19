//切片append,copy
package main
import "fmt"

func InsertStringSliceCopy(slice,insertion []string, index int) []string {
    result := make([]string, len(slice)+len(insertion)) 
    at := copy(result, slice[:index])
    at += copy(result[at:], insertion)
    copy(result[at:], slice[index:])
    return result
}


func main() {
    s := []string {"M", "N", "O", "P", "Q", "R"}
    x := InsertStringSliceCopy(s,[]string{"a","b","c"}, 0) //at the front
    y := InsertStringSliceCopy(s,[]string{"x","y","z"}, 3) //in the middle
    z := InsertStringSliceCopy(s,[]string{"z"}, len(s)) //at the end
    fmt.Printf("%v\n%v\n%v\n%v\n", s,x,y,z)
}
