package main
import "fmt"

type PathError struct {
    Op string
    Path string
    Err error
}

func (e *PathError) Error() string{
    return e.Op + " " + e.Path + ": " +e.Err.Error()
}

func Stat(name string) (fi FileInfo, err error) {
    var stat syscall.Stat_t
    err = syscall.Stat(name, &stat)
    if err != nil {
	return nil, &PathError{"stat", name, err}
    }

    return fileInfoFromStat(&stat, name),nil
}

func main(){
    fi, err := os.Stat("a.txt")
    if err != nil {
	if e, ok := err.(*os.PathError); ok && e.Err != nil {
	    //
	    fmt.Println(e, ok)
	}
    }
}
