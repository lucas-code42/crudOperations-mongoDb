package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lucas-code42/testMongoDb/pkg/api/utils"
)

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("could not create book"))
		return
	}

	book, err := utils.ParseToModel(requestBody)
	if err != nil {
		w.WriteHeader(422)
		w.Write([]byte("could not create book"))
		return
	}

	if err = book.CheckFields(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	responseDto := c.service.CreateBookService(book)
	if responseDto.Error != nil {
		json.NewEncoder(w).Encode(responseDto)
	}
	json.NewEncoder(w).Encode(responseDto)
}
