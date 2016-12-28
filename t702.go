//子网掩码地址
//编译 go build t702
//使用 ./t702 192.168.0.1

package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage:%s ip.addr\n", os.Args[0])
		os.Exit(1)
	}

	dotaddr := os.Args[1]
	addr := net.ParseIP(dotaddr)
	if addr ==nil {
		fmt.Println("Invalid address")
	}

	mask := addr.DefaultMask()
	fmt.Println("Subnet mask is: ", mask.String())
	network := addr.Mask(mask)
	fmt.Println("Network address is；", network.String())
	ones,bits := mask.Size()
	fmt.Println("Mask bits:", ones, " Total bits:", bits)
	os.Exit(0)


}
