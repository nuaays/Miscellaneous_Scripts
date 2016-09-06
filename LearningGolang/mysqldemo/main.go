package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	testConnection()
}

func testConnection() {
	//	change	database	user	and	password
	db, err := sql.Open("mysql", string("root:test@tcp(10.139.99.150:3306)/auto_test_00"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()        //	test	connection
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected")
	defer db.Close()
}


func insertData(){
	//db, err = sql.Open("mysql", string("root:test@tcp(10.139.99.150:3306)/auto_test_00"))
	//if err != nil {
	//	panic(err)
	//}
	//defer	db.Close()
	//err = db.Ping() // test connection
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println("connected")
	//
	////prepare development


}