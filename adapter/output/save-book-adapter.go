package output

import (
	"database/sql"
	"fmt"
	"go-clean-code/entity"
)

type SaveBookAdapter struct {
	db *sql.DB
}

func NewSaveBookAdapter(db *sql.DB) *SaveBookAdapter {
	return &SaveBookAdapter{db}

}

func (s *SaveBookAdapter) Save(book entity.Book) entity.Book {
	sqlStatement := fmt.Sprintf("INSERT INTO public.book VALUES ($1,$2)")
	insert, err := s.db.Prepare(sqlStatement)
	checkErr(err)
	result, err := insert.Exec(book.Name, book.Loaned)
	checkErr(err)
	affect, err := result.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	return book
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
