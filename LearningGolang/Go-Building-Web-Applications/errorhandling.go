

import (
  "fmt"
  "net/http"
  "time"
)

const (
  Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
  response := "The time is now " + time.Now().String()
  fmt.Fprintln(w,response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static.html")
}


func serveError() {
  fmt.Println("There's no way I'll work!")
}

func main() {
  http.HandleFunc("/static", serveStatic)
  http.HandleFunc("/", serveDynamic)
  http.HandleFunc("/error",serveError)
  http.ListenAndServe(Port, nil)
}

