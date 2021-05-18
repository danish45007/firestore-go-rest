package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	port := ":8080"

	// root handler
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello-World")
	})
	// get post handler
	router.HandleFunc("/get-post", getPosts).Methods("GET")
	// create post habndler
	router.HandleFunc("/create-post", createPost).Methods("POST")
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe("127.0.0.1:8080", router))
}
