package book

import (
	"go-clean-code/entity"
)

type SaveBookService struct {
}

func NewSaveBookService() *SaveBookService {
	return &SaveBookService{}
}

func (s *SaveBookService) Save(name string, channel chan entity.Book) {
	go func() {
		channel <- entity.Book{Name: name}
	}()
}
