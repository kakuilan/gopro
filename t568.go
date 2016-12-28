//bytes.SplitN
package main
import (
		"fmt"
		"bytes"
		)

func main(){
	s := []byte("南瓜,黄瓜,西红柿,茄子")
	sep := []byte(",")
	n := 2
	s2 := bytes.SplitN(s,sep,n)

	for _,f := range s2 {
		fmt.Printf("%q\n",f)
	}
	fmt.Println("s2: ", s2)

	s3 := bytes.SplitN(s,sep,0)
	for _,f := range s3 {
		fmt.Printf("%q\n", f)
	}
	fmt.Println("s3: ", s3)

	s4 := bytes.SplitN(s,sep, -1)
	for _,f := range s4 {
		fmt.Printf("%q\n", f)
	}
	fmt.Println("s4: ", s4)


}

