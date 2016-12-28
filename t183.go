// mysql 批量重复插入
package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "math/rand"
    "time"
)

func main(){
    db,err := sql.Open("mysql","test:123456@tcp(192.168.128.1:3306)/test?charset=utf8")
    if err != nil {
	fmt.Println(err)
	return
    }
    defer db.Close()
    var smt *sql.Stmt
    smt,err = db.Prepare("insert into person(name,age,IsBoy) values(?,?,?)")
    if err != nil {
	fmt.Println(err)
	return
    }
    fmt.Println("开始插入数据...", time.Now())

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i:=0;i<10000;i++ {
	_,err = smt.Exec(fmt.Sprintf("张%d", r.Int()), r.Intn(50), r.Intn(100)%2)
	if err != nil {
	    fmt.Println(err)
	    return
	}
    }
    fmt.Println("数据插入完成!", time.Now())



}
