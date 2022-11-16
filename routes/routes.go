package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/adapter/input"
	"go-clean-code/service"
)

type Routes struct {
	server *gin.Engine
}

func NewRoutes(server *gin.Engine) {
	r := Routes{server: server}

	v1 := r.server.Group("/v1")
	{
		book := v1.Group("/book")
		{
			var (
				bookService = service.NewSaveBookService()
				adapter     = input.NewSaveBookAdapter(bookService)
			)
			book.POST("", adapter.CreateBook())
		}
	}
}
