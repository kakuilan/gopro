//break可用于for,switch,select; 而continue仅能用于for
package main

func main(){
	x := 100
	switch{
		case x>= 0:
			if x==0 {break}
			println(x)
	}


}
