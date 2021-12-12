package handlers

import (
	"crud_with_rest_api_on_mock_data/mocks"
	"crud_with_rest_api_on_mock_data/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//read dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	//read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var updateBook models.Book
	json.Unmarshal(body, &updateBook)

	//search for id and update

	for index, book := range mocks.Books {
		if book.Id == id {
			book.Title = updateBook.Title
			book.Author = updateBook.Author
			book.Desc = updateBook.Desc

			mocks.Books[index] = book
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("updated")
			break
		}
	}
}
