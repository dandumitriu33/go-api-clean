// Go-API-clean is a clean architecture simple GO API with Mux.
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	const port string = ":8000"
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running")
	})
	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
