package repository

import (
	"context"
	"log"

	"github.com/lucas-code42/testMongoDb/pkg/api/config"
	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
)

func (br *MongoBookRepository) CreateBook(book entity.Book) error {
	coll := br.DataBaseConnection.Collection(
		config.GetMongoConfig().DefaultCollection,
	)

	result, err := coll.InsertOne(context.Background(), book)
	if err != nil {
		return err
	}

	log.Println("inserted ID:", result.InsertedID)
	return nil
}
