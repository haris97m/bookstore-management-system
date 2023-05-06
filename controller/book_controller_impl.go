package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/haris97m/bookstore-management-system/model"
)

type BookStore struct {
	books []model.Book
	file  string
}

func NewBookStore() *BookStore {
	return &BookStore{
		file: "Book.json",
	}
}

// Load JSON file
func (bookStore *BookStore) load() error {
	data, _ := ioutil.ReadFile(bookStore.file)

	json.Unmarshal(data, &bookStore.books)

	return nil
}

// Save JSON file.
func (bookStore *BookStore) save() error {
	data, _ := json.Marshal(bookStore.books)

	ioutil.WriteFile(bookStore.file, data, 0644)

	return nil
}

func (bookStore *BookStore) AddBook(book model.Book) {
	bookStore.load()
	bookStore.books = append(bookStore.books, book)
	bookStore.save()
	fmt.Println("Success add Book:", book)
}

func (bookStore *BookStore) GetBookByID(ID int) model.Book {
	bookStore.load()
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
	bookStore.load()
	for i, book := range bookStore.books {
		if book.ID == ID {
			bookStore.books = append(bookStore.books[:i], bookStore.books[i+1:]...)
			bookStore.save()
			fmt.Println("Success remove Book with ID:", ID)
			return
		}
	}
}
