package controller

import (
	"testing"

	"github.com/haris97m/bookstore-management-system/model"
)

func TestBookHandler_AddBook(t *testing.T) {
	bookstore := NewBookStore()
	newBook := model.Book{ID: 1, Title: "One Piece", Author: "Eiichiro Oda", Quantity: 1_000, Available: true}
	bookstore.AddBook(newBook)

	if len(bookstore.books) != 1 {
		t.Errorf("Expected added 1 book, but got %v", len(bookstore.books))
	}

	bookstore.RemoveBookByID(1)
}

func TestBookHandler_GetBookByID(t *testing.T) {
	bookstore := NewBookStore()
	newBook := model.Book{ID: 1, Title: "One Piece", Author: "Eiichiro Oda", Quantity: 100, Available: true}
	bookstore.AddBook(newBook)

	// Test Found Book
	foundBook := bookstore.GetBookByID(1)
	if foundBook.ID != newBook.ID {
		t.Errorf("Expected book with ID %v, but got book with ID %v", newBook.ID, foundBook.ID)
	}

	// Test Not Found Book
	notFoundBook := bookstore.GetBookByID(2)
	if notFoundBook.ID != 0 {
		t.Errorf("Expected to not find book, but got book with ID %v", notFoundBook.ID)
	}

	bookstore.RemoveBookByID(1)
}

func TestBookHandler_RemoveBookByID(t *testing.T) {
	bookstore := NewBookStore()
	newBook := model.Book{ID: 1, Title: "One Piece", Author: "Eiichiro Oda", Quantity: 100, Available: true}
	bookstore.AddBook(newBook)

	bookstore.RemoveBookByID(1)

	if len(bookstore.books) != 0 {
		t.Errorf("Expected 1 book, but got %v", len(bookstore.books))
	}

	notFoundBook := bookstore.GetBookByID(1)
	if notFoundBook.ID != 0 {
		t.Errorf("Expected to not find book, but got book with ID %v", notFoundBook.ID)
	}
}
