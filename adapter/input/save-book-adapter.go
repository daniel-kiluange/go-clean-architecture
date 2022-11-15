package input

import (
	"go-clean-code/entity"
	"go-clean-code/useCase/input"
)

type SaveBookAdapter struct {
	useCase input.SaveBookUseCase
}

func NewSaveBookAdapter(useCase input.SaveBookUseCase) *SaveBookAdapter {
	return &SaveBookAdapter{useCase: useCase}
}

func (a *SaveBookAdapter) CreateBook(name string, channel chan entity.Book) {

	go func() {
		channel <- a.useCase.Save(name)
	}()
}
