//bytes.SplitAfterN
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("managerteacherworkerframerstudent")
	sep := []byte("er")
	n := 3
	s2 := bytes.SplitAfterN(s,sep,n)
	for _,f := range s2 {
		fmt.Printf("%q\n",f)
	}
	fmt.Println(s2)

	s3 := bytes.SplitAfterN(s,sep,0)
	for _,f := range s3{
		fmt.Printf("%q\n",f)
	}
	fmt.Println(s3)

	s4 := bytes.SplitAfterN(s,sep,-1)
	for _,f := range s4{
		fmt.Printf("%q\n",f)
	}
	fmt.Println(s4)


}
