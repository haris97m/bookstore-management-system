package controller

import (
	"fmt"

	"github.com/haris97m/bookstore-management-system/model"
)

type BookStore struct {
	books []model.Book
}

func (bookStore *BookStore) AddBook(book model.Book) {
	bookStore.books = append(bookStore.books, book)
	fmt.Println("Success add Book:", book)
}

func (bookStore *BookStore) GetBookByID(ID int) model.Book {
	for _, book := range bookStore.books {
		if book.ID == ID {
			fmt.Println(book)
			return book
		}
	}
	fmt.Println("Not found Book with ID:", ID)
	return model.Book{}
}

func (bookStore *BookStore) RemoveBookByID(ID int) {
	for i, book := range bookStore.books {
		if book.ID == ID {
			bookStore.books = append(bookStore.books[:i], bookStore.books[i+1:]...)
			fmt.Println("Success remove Book with ID:", ID)
		}
	}
}
