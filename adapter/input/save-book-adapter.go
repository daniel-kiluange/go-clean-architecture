package input

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/entity"
	"go-clean-code/useCase/input"
	"net/http"
)

type SaveBookAdapter struct {
	useCase input.SaveBookUseCase
}

func NewSaveBookAdapter(useCase input.SaveBookUseCase) *SaveBookAdapter {
	return &SaveBookAdapter{useCase: useCase}
}

func (a *SaveBookAdapter) CreateBook() gin.HandlerFunc {
	var bookChannel = make(chan entity.Book, 1)
	return func(context *gin.Context) {
		var book entity.Book
		err := context.BindJSON(&book)
		if err != nil {
			return
		}
		a.useCase.Save(book.Name, bookChannel)
		context.JSON(http.StatusCreated, <-bookChannel)
	}
}
