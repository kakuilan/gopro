//bytes.SplitAfetr
package main
import (
		"fmt"
		"bytes"
		)


func main(){
	s := []byte("managerteacherworkerfarmerstudent")
	sep := []byte("er")
	s2 := bytes.SplitAfter(s,sep)

	for _,f := range s2 {
		fmt.Printf("%q\n",f)
	}
	fmt.Println(s2)



}
