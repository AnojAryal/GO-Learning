package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port = ":8000"

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Up and Running...")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Up and Running..."))
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
