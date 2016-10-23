package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

const (
  PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  pageID := vars["id"]
  fileName := "files/" + pageID + ".html"
  http.ServeFile(w,r,fileName)
}


func main() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/pages/{id:[0-9]+}",pageHandler)
  http.Handle("/",rtr)
  http.ListenAndServe(PORT,nil)
}
