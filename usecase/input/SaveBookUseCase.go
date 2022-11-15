package input

import "go-clean-code/entity"

type SaveBookUseCase interface {
	Save(name string) entity.Book
}
