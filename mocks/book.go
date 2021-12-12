package mocks

import "crud_with_rest_api_on_mock_data/models"

var Books = []models.Book{
	{Id: 1,
		Title:  "golang",
		Author: "Jeewan Sharma",
		Desc:   "book to learn golang"},
	{Id: 2,
		Title:  "golang2",
		Author: "Jeewan Sharma",
		Desc:   "book to learn golang2.0"},
}
