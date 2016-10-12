package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type API struct {
	Message string "json:message"
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{"Hello, world!"}
		output, err := json.Marshal(message)

		if err != nil {
			fmt.Println("Something went wrong!")
		}

		fmt.Fprintf(w, string(output))
	})

	http.HandleFunc("/api/user/\d+", func(w http.ResponseWriter, r *http.Request) {
	// react dynamically to an ID as supplied in the URL

	})


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
