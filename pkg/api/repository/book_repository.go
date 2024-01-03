package repository

import (
	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepositoryImp interface {
	GetBook() []entity.Book
	CreateBook(book entity.Book) error
	UpdateBook(book entity.Book) error
	DeleteBook(book entity.Book) error
}

type BookRepository struct {
	DataBaseConnection *mongo.Client
}

// todo: mongo client não deveria ser diretamente a tabela que vamos usar?
func NewBookRepository(dataBaseConnection *mongo.Client) BookRepositoryImp {
	return &BookRepository{
		DataBaseConnection: dataBaseConnection,
	}
}
