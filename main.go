package main

import (
	"fmt"

	"github.com/haris97m/bookstore-management-system/controller"
	"github.com/haris97m/bookstore-management-system/model"
)

func main() {
	bookstore := controller.NewBookStore()

	fmt.Println()
	fmt.Println("Bookstore Management System")
	fmt.Println("===========================")
	fmt.Println()

	// AddBook
	fmt.Println("=>ADD BOOK")
	bookstore.AddBook(model.Book{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 10, Available: true})
	bookstore.AddBook(model.Book{ID: 2, Title: "1984", Author: "George Orwell", Quantity: 0, Available: false})
	bookstore.AddBook(model.Book{ID: 3, Title: "Naruto", Author: "Masashi Kishimoto", Quantity: 2, Available: true})
	fmt.Println()

	// GetBookByID
	fmt.Println("=>GET BOOK BY ID")
	bookstore.GetBookByID(1)
	bookstore.GetBookByID(2)
	bookstore.GetBookByID(3)
	fmt.Println()

	// RemoveBookByID
	fmt.Println("=>REMOVE BOOK BY ID")
	bookstore.RemoveBookByID(1)
	fmt.Println()

	fmt.Println("=>GET BOOK AFTER DELETED")
	bookstore.GetBookByID(1)
	fmt.Println()

}
