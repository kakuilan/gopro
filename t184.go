//mysql 事务
package main
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main(){
    db,err := sql.Open("mysql","test:123456@tcp(192.168.128.1:3306)/test?charset=utf8")
    if err != nil {
	fmt.Println(err)
	return
    }

    defer db.Close()
    var trans *sql.Tx
    trans,err = db.Begin()
    if err != nil {
	fmt.Println(err)
	return
    }

    _,err = trans.Exec("insert intoperson(name,age,IsBoy) values('zhang 3',99,false)")
    if err != nil {
	trans.Rollback()
	fmt.Println(err)
    }else{
	trans.Commit()
    }

}
