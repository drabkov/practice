package main

import (
	"github.com/gorilla/mux"
	"log"
	"myserver/restapi/handlers"
	"net/http"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/{id}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/{id}", handlers.UpdateBookByID).Methods("PUT")
	r.HandleFunc("/{id}", handlers.DeleteBookByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}
