//sort.Search二分查找
package main
import (
    "fmt"
    "sort"
    "strings"
)

func main() {
    files := []string {"Test.conf","util.go","Makefile","misc.go","main.go"}
    target := "Makefile"

    //普通遍历查找
    fmt.Printf("%q\n", files)
    for i,file := range files {
        if file == target {
            fmt.Printf("found \"%s\" at files[%d]\n", file, i) 
            break
        }
    }
    fmt.Println()

    //sort.Search二分搜索算法
    sort.Strings(files)
    fmt.Printf("%q\n", files)
    i := sort.Search(len(files), func(i int) bool {return files[i] >= target})
    if i<len(files) && files[i] == target {
        fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
    }

    //搜索一个小写的
    target = "makefile"
    fmt.Println()
    caseInsensitiveCompare := func(i int) bool {
        return  strings.ToLower(files[i]) >= target
    }
    i = sort.Search(len(files), caseInsensitiveCompare)
    if i<=len(files) && strings.EqualFold(files[i], target) {
        fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
    }



}

