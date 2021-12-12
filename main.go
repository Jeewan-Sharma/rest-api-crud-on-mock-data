package main

import (
	"crud_with_rest_api_on_mock_data/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET") //http.MethodGet
	router.HandleFunc("/books", handlers.AddBooks).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")       //http.MethodGet
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")    //http.MethodPut
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE") //http.MethodPut

	log.Println("api is running")
	http.ListenAndServe(":4000", router)
}
