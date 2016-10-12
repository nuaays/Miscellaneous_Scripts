package main

import (
	"encoding/json"
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {

	urlParams := r.URL.Query()
	name := urlParams.Get(":name")
	HelloMessage := "Hello, " + name

	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong!")
	}

	fmt.Fprintf(w, string(output))

}

func main() {
	mux := routes.New()
	mux.Get("/api/:name", Hello)
	http.Handle("/", mux)

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
