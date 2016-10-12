package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)


type User struct {
  ID int "json:id"
  Name  string "json:username"
  Email string "json:email"
  First string "json:first"
  Last  string "json:last"
}



func CreateUser(w http.ResponseWriter, r *http.Request) {

  NewUser := User{}
  NewUser.Name = r.FormValue("user")
  NewUser.Email = r.FormValue("email")
  NewUser.First = r.FormValue("first")
  NewUser.Last = r.FormValue("last")
  output, err := json.Marshal(NewUser)
  fmt.Println(string(output))
  if err != nil {
    fmt.Println("Something went wrong!")
  }

  sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
  q, err := database.Exec(sql)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(q)
}

routes := mux.NewRouter()
routes.HandleFunc("/api/user/create", CreateUser).Methods("GET")
