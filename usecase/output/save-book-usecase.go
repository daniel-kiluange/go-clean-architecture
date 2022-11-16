package output

import "go-clean-code/entity"

type SaveBookUseCase interface {
	Save(entity.Book) entity.Book
}
