package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/handlers"
	
	"myserver/restapi"	
)

func main() {
	router := restapi.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// launch server
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}