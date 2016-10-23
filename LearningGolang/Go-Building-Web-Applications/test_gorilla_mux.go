package main

import (
  "os"
  "github.com/gorilla/mux"
  "net/http"
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "log"
)

const (
  DBHost  = "127.0.0.1"
  DBPort  = ":3306"
  DBUser  = "root"
  DBPass  = "password!"
  DBDbase = "cms"
)

const (
  PORT = ":8080"
)

var database *sql.DB

type Page struct {
  Title   string
  Content string
  Date    string
}


func pageHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  pageID := vars["id"]
  fileName := "files/" + pageID + ".html"
  err := os.Stat(fileName)
  if err != nil{
    fileName = "files/404.html"
  }
  http.ServeFile(w,r,fileName)
}


func main() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/pages/{id:[0-9]+}",pageHandler)
  http.Handle("/",rtr)
  http.ListenAndServe(PORT,nil)
}
