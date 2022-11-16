package input

import "go-clean-code/entity"

type CreateBookUseCase interface {
	Save(name string, channel chan entity.Book)
}
