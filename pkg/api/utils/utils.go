package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
)

func ParseToModel(data []byte) (entity.Book, error) {
	var book entity.Book
	
	if err := json.Unmarshal(data, &book); err != nil {
		return entity.Book{}, fmt.Errorf("error to parse")
	}

	log.Println("Função ParseToModel", book)
	return book, nil
}

func ParseToBytes(data entity.Book) ([]byte, error) {
	return json.Marshal(data)
}
