package input

import "go-clean-code/entity"

type SaveBookUseCase interface {
	Save(name string, channel chan entity.Book)
}
