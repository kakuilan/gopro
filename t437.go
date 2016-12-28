//for
package main
import "fmt"

func main(){
	sum := 0;
	for index:=0;index<10;index++{
		sum += index;
	}
	fmt.Println("sum is equal to ", sum)

	//简化
	s2 := 1
	for ;s2<100; {
		s2 +=s2
	}
	fmt.Println("s2=", s2)

	//再简化
	s3 :=1
	for s3<100{
		s3 += s3
	}
	fmt.Println("s3=", s3)


}
