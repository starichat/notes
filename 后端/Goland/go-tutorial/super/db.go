package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	db, err := sql.Open("mysql", "root:root@/wxl?charset=utf8")
	checkErr(err)
	defer func() {
		fmt.Println("close")
		db.Close()
	}()
	//insert data
	stmt, err := db.Prepare("Insert userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("wxl", "go", "2019-08-20")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("Update userinfo set departname=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("go develop", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	rows, err := db.Query("Select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}
	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	fmt.Println(affect)

}
