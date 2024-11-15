package store

import "github.com/google/uuid"

var books = []*Book{
	{
		ID:     "1",
		UserID: "1",
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
	},
	{
		ID:     "2",
		UserID: "1",
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
	},
	{
		ID:     "3",
		UserID: "2",
		Title:  "1984",
		Author: "George Orwell",
	},
	{
		ID:     "4",
		UserID: "2",
		Title:  "Animal Farm",
		Author: "George Orwell",
	},
}

type Book struct {
	ID     string
	UserID string
	Title  string
	Author string
}

func Books() []*Book {
	return books

}

func CreateBook(author string, title string, userID string) *Book {
	book := &Book{
		ID:     uuid.New().String(),
		UserID: userID,
		Title:  title,
		Author: author,
	}
	books = append(books, book)
	return book
}
