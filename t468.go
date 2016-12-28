//mysql数据库使用
package main
import (
		_ "github.com/go-sql-driver/mysql"
		"database/sql"
		"fmt"
		)

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}


func main(){
	db,err := sql.Open("mysql", "root:root12@tcp(192.168.128.1:3306)/test?charset=utf8")
	//fmt.Println(db, err)
	checkErr(err)

	//插入数据
	stmt,err := db.Prepare("INSERT person SET name=?,age=?,IsBoy=?")
	checkErr(err)

	res,err := stmt.Exec("abcx", 24, 1)
	checkErr(err)
	
	id,err := res.LastInsertId()
	if err == nil {
		fmt.Println("插入成功,ID是", id)
	}

	//更新数据
	stmt,err = db.Prepare("UPDATE person SET name=? WHERE id=?")
	checkErr(err)

	res,err = stmt.Exec("ttu", id)
	checkErr(err)

	affect,err := res.RowsAffected()
	checkErr(err)
	if err == nil {
		fmt.Println("更新成功,", affect)
	}

	//查询数据
	rows,err := db.Query("SELECT * FROM person")
	checkErr(err)

	for rows.Next(){
		var uid int
		var uname string
		var age int
		var sex int
		err = rows.Scan(&uid, &uname, &age, &sex)
		checkErr(err)

		fmt.Println("信息 名称:",uid, " 年龄:",age, " 性别:", sex)
	}

	//删除数据
	stmt,err = db.Prepare("DELETE FROM person WHERE id=?")
	checkErr(err)

	res,err = stmt.Exec(id)
	checkErr(err)

	affect,err = res.RowsAffected()
	checkErr(err)

	fmt.Println("删除结果:", affect)

	db.Close()


}

