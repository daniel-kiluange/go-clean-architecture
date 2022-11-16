package entity

import "fmt"

type Book struct {
	Name   string `json:"name"`
	Loaned bool   `json:"lent"`
}

func (b *Book) Lent(i int) {
	fmt.Println("Book", b.Name, "is borrowed by user", i)
	b.Loaned = true
}
