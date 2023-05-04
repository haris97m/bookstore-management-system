package controller

import (
	"github.com/haris97m/bookstore-management-system/model"
)

type BookHandler interface {
	AddBook(book model.Book)
	GetBookByID(ID int) model.Book
	RemoveBookByID(ID int)
}
