package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var database *sql.DB

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID    int    "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
	Image string "json:image"
}

type CreateResponse struct {
	Error     string "json:error"
	ErrorCode int    "json:code"
}

type ErrMsg struct {
	ErrCode    int
	StatusCode int
	Msg        string
}

func ErrorMessages(err int64) ErrMsg {
	var em ErrMsg
	errorMessage := ""
	statusCode := 200
	errorCode := 0
	switch err {
	case 1062:
		errorMessage = "Duplicate entry"
		errorCode = 10
		statusCode = 409
	}

	em.ErrCode = errorCode
	em.StatusCode = statusCode
	em.Msg = errorMessage

	return em
}

func dbErrorParse(err string) (string, int64) {
	Parts := strings.Split(err, ":")
	errorMessage := Parts[1]
	Code := strings.Split(Parts[0], "Error ")
	errorCode, _ := strconv.ParseInt(Code[1], 10, 32)
	return errorMessage, errorCode
}

func userRouter(w http.ResponseWriter, r *http.Request) {
	ourUser := User{}
	ourUser.Name = "Sen Yang"
	ourUser.Email = "yangsen@zhongan.com"
	ourUser.ID = 100
	//ourUser.First = "Sen"
	// ourUser.Last = "Yang"

	output, _ := json.Marshal(&ourUser)
	fmt.Fprintln(w, string(output))

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	NewUser.Image = ""
	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	// database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()

	Response := CreateResponse{}

	sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "', user_image='" + NewUser.Image + "'"
	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
		Response.Error = err.Error()
	}
	fmt.Println(q)
	createOutput, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(createOutput))

}

func UsersUpdate(w http.ResponseWriter, r *http.Request) {
	Response := CreateResponse{}
	params := mux.Vars(r)
	uid := params["id"]
	email := r.FormValue("email")
	var userCount int
	err := database.QueryRow("SELECT COUNT(user_id) FROM users WHERE user_id=?", uid).Scan(&userCount)
	if userCount == 0 {
		//error, httpCode, msg := ErrorMessages(404)
		em := ErrorMessages(404)
		log.Println(em.ErrCode)
		log.Println(w, em.Msg, em.StatusCode)

	} else if err != nil {
		log.Println(err)
	} else {
		//update
		_, uperr := database.Exec("UPDATE users SET user_email=?WHERE user_id=?", email, uid)
		if uperr != nil {
			//_, errorCode := dbErrorParse(uperr.Error())
			//_, httpCode, msg := ErrorMessages(errorCode)

		} else {
			Response.Error = "success"
			Response.ErrorCode = 0
			output, _ := json.Marshal(Response)
			fmt.Println(w, string(output))
		}

	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	ReadUser := User{}

	// database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()

	w.Header().Set("Pragma", "no-cache")

	query_err := database.QueryRow("select * from users where user_id=?", id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email, &ReadUser.Image)
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

func UsersRetrieve(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Pragma", "no-cache")
	// database, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer database.Close()

	rows, _ := database.Query("select * from users LIMIT 10")
	Response := Users{}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email, &user.Image)

		o, _ := json.Marshal(user)
		fmt.Println(string(o))

		Response.Users = append(Response.Users, user)
	}
	output, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

func main() {

	var apiversion string = "1.0"
	fmt.Println("Starting JSON server")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/social_network")
	if err != nil {
		log.Fatal(err)
		errorMessage, errorCode := dbErrorParse(err.Error())
		fmt.Println(errorMessage)
		fmt.Println(errorCode)
		// error, httpCode, msg := ErrorMessages(errorCode)
		// Response.Error = msg
		// Response.ErrorCode = error
		// fmt.Println(httpCode)
		//http.Error(w, "Conflict", httpCode)
	}
	//defer db.Close()
	database = db
	fmt.Println("Waiting ...")

	gorillaRoute := mux.NewRouter()
	//gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	//gorillaRoute.HandleFunc("/api/{user:\\w+}", Hello)

	//create user
	//http://localhost:8080/api/user/create?user=nkozyra&first=Nathan&last=Kozyra&email=nathan@nathankozyra.com
	gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
	//get user information
	gorillaRoute.HandleFunc("/api/user/read/{id:\\d+}", GetUser).Methods("GET")
	//userupdate
	gorillaRoute.HandleFunc("/api/users/{id:[0-9]+}", UsersUpdate).Methods("PUT")
	//user retrieve
	gorillaRoute.HandleFunc("/api/users", UsersRetrieve).Methods("GET")

	//userRouter
	gorillaRoute.HandleFunc(fmt.Sprintf("/api/%s/user", apiversion), userRouter).Methods("GET")

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
