package main

import (
	_ "go-learning/docs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Your API Title
// @version 1.0
// @description This is a sample server for your Go application.
// @host localhost:8000
// @BasePath /

func main() {

	db, err := InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("Up and Running...")
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Up and Running..."))
	})

	router.HandleFunc("/posts", getPosts(db)).Methods("GET")
	router.HandleFunc("/posts", addPost(db)).Methods("POST")

	// Serve Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	const port = ":8000"
	log.Println("Server listening on Port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
