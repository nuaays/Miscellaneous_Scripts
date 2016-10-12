package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	name := urlParams["user"]
	HelloMessage := "Hello, " + name

	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	//gorillaRoute.HandleFunc("/api/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/{user:\\w+}", Hello)
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
