package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID    int    "json:id"
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

	database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'"
	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Pragma", "no-cache")
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}

	database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	query_err := database.QueryRow("select * from users where user_id=?", id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)
	switch {
	case query_err == sql.ErrNoRows:
		fmt.Fprintf(w, "No such user")
	case query_err != nil:
		log.Fatal(query_err)
		fmt.Fprintf(w, "Error")
	default:
		output, _ := json.Marshal(ReadUser)
		fmt.Fprintf(w, string(output))
	}

}

func main() {
	gorillaRoute := mux.NewRouter()
	//gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	//gorillaRoute.HandleFunc("/api/{user:\\w+}", Hello)

	//create user
	//http://localhost:8080/api/user/create?user=nkozyra&first=Nathan&last=Kozyra&email=nathan@nathankozyra.com
	gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	gorillaRoute.HandleFunc("/api/user/read/{id:\\d+}", GetUser).Methods("GET")
	http.Handle("/", gorillaRoute)

	// Server := Server {
	//   Addr: ":8080",
	//   Handler: urlHandler,
	//   ReadTimeout: 1000 * time.MicroSecond,
	//   WriteTimeout: 1000 * time.MicroSecond,
	//   MaxHeaderBytes: 0,
	//   TLSConfig: nil
	// }
	http.ListenAndServe(":8080", nil)
}
