//defer,后进先出
package main

func main(){
	for i:=0;i<5;i++{
		defer println("defer ", i)
	}



}
