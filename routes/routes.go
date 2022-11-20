package routes

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/adapter/input"
	"go-clean-code/service/book"
)

type Routes struct {
	server *gin.Engine
}

func NewRoutes(server *gin.Engine) {
	bookService := book.NewSaveBookService()
	input.NewSaveBookAdapter(bookService).WithServer(server).Serve()
}
