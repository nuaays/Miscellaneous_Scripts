package main


import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
)


func main() {
    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dbname?charset=utf8")
    if err != nil {
       panic(err.Error())
    }
    defer db.Close()
   
    rows, err := db.Query("select * from users")
    if err != nil {
       fmt.Println(err.Error())
    }
    defer rows.Close()

    var id int
    var name string
    for rows.Next() {
        r_err := rows.Scan(&id, &name);
    }
}
