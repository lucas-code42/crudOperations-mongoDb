package service

import (
	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
	"github.com/lucas-code42/testMongoDb/pkg/api/model/response"
	"github.com/lucas-code42/testMongoDb/pkg/api/repository"
)

type ServiceImp interface {
	GetBookService() response.ResponseDto
	CreateBookService(book entity.Book) response.ResponseDto
	UpdateBookService(book entity.Book) response.ResponseDto
	DeleteBookService(book entity.Book) response.ResponseDto
}

type BookService struct {
	bookRepo repository.BookRepositoryImp
}

func NewBookService(bookRepository repository.BookRepositoryImp) ServiceImp {
	return &BookService{
		bookRepo: bookRepository,
	}
}
