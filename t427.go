//error类型
package main
import (
		"fmt"
		"errors"
		)

func main(){
	err := errors.New("emit macho dear:sdfaafd")
	if err != nil {
		fmt.Println(err)
	}
}
