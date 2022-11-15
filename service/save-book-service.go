package service

import (
	"go-clean-code/entity"
)

type SaveBookService struct {
}

func NewSaveBookService() *SaveBookService {
	return &SaveBookService{}
}

func (s *SaveBookService) Save(name string) entity.Book {
	return entity.Book{Name: name}
}
