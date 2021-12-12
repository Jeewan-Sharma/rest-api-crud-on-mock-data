package handlers

import (
	"crud_with_rest_api_on_mock_data/mocks"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	//reading dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//searching ID
	for _, book := range mocks.Books {
		if book.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
}
