package entity

import "fmt"

type Book struct {
	Name     string
	isBorrow bool
}

func (b *Book) Borrow(i int) {
	fmt.Println("Book", b.Name, "is borrowed by user", i)
	b.isBorrow = true
}
