package model

import (
	"github.com/go-framework/web"
)

type Book struct {
	Id       int
	UserName string
	Content  string
}

func (b *Book) Insert() bool {
	db := web.DB.Open()
	defer db.Close()
	sql := "insert into books(username, content) values(?,?)"
	_, err := db.Exec(sql, b.UserName, b.Content)
	if err != nil {
		return false
	}
	return true
}

func (b *Book) GetAll() []*Book {
	books := make([]*Book, 0)
	db := web.DB.Open()
	defer db.Close()
	sql := "SELECT id,username,content FROM books"
	rows, err := db.Query(sql)
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		book := &Book{}
		err := rows.Scan(&book.Id, &book.UserName, &book.Content)
		if err != nil {
			return nil
		}
		books = append(books, book)
	}
	return books
}
