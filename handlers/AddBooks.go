package handlers

import (
	"crud_with_rest_api_on_mock_data/mocks"
	"crud_with_rest_api_on_mock_data/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func AddBooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var book models.Book
	json.Unmarshal(body, &book)

	//appending the mocks data
	book.Id = rand.Intn(100)
	mocks.Books = append(mocks.Books, book)

	//
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("created")
}
