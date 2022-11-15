package main

import (
	"fmt"
	"go-clean-code/adapter/input"
	"go-clean-code/entity"
	"go-clean-code/service"
)

var (
	bookService = service.NewSaveBookService()
	adpter      = input.NewSaveBookAdapter(bookService)
)

func main() {
	bookChannel := make(chan entity.Book, 1)

	adpter.CreateBook("Daniel", bookChannel)

	book := <-bookChannel

	fmt.Println(book)

	book.Borrow(1)

	fmt.Println(book)
}
