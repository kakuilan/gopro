//strings包
package main

import (
	"fmt"
	"strings"
)

func main() {
	names := "Niccolo·Geoffrey·Amelie··Turlought·Jose"
	fmt.Print("|")
	for _, name := range strings.Split(names, "·") {
		fmt.Printf("%s|", name)
	}

	fmt.Println()

    for _,record := range []string{"Laszlo Lajtha*1892*1963", "Edouard Lalo\t1823\t1892", "Jose Angel Mamas|1775|1814"} {
        fmt.Println(strings.FieldsFunc(record, func(char rune) bool {
            switch char {
                case '\t','*','|':
                    return true
            }
            return false
        }))
    }

}
