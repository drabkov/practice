package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"myserver/domain"
	"myserver/infrastructure"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := infrastructure.GetBooks()
	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(books) // encode similar to serialize process.
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)
	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		getError(err, w)
		return
	}
	book, err := infrastructure.GetBookByID(id)
	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func GetBookByNamePublishingHouse(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	// we get params with mux.
	var params = mux.Vars(r)

	book, err := infrastructure.GetBookByNamePublishingHouse(params["name"])
	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book domain.Book

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	// insert our book model.
	fmt.Println(r.Body)
	fmt.Println(&book)
	result, err := infrastructure.CreateBook(&book)

	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		getError(err, w)
		return
	}

	var book domain.Book
	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	err = infrastructure.UpdateBookByID(id, &book)
	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		getError(err, w)
		return
	}

	deleteResult, err := infrastructure.DeleteBookByID(id)
	if err != nil {
		getError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
func getError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
