package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lucas-code42/testMongoDb/pkg/api/model/entity"
	"github.com/lucas-code42/testMongoDb/pkg/api/model/response"
)

func (bs *BookService) CreateBookService(book entity.Book) response.ResponseDto {
	// aplicar logica de domínio...
	book.Id = fmt.Sprintf("%s-%s", "BookDomain", uuid.NewString())

	if err := bs.bookRepo.CreateBook(book); err != nil {
		return response.ResponseDto{
			Error: err,
		}
	}
	
	return response.ResponseDto{
		Data: []map[string]any{
			{"created": book},
		},
		Error: nil,
	}
}
