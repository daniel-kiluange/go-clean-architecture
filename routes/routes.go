package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/adapter/input"
	bookServiceImpl "go-clean-code/service/book"
)

type Routes struct {
	server *gin.Engine
}

func NewRoutes(server *gin.Engine) {
	r := Routes{server: server}

	book := r.server.Group("/book")
	{
		var (
			bookService = bookServiceImpl.NewSaveBookService()
			adapter     = input.NewSaveBookAdapter(bookService)
		)
		book.POST("", adapter.CreateBook())
	}
}
