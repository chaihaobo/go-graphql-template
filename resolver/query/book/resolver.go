package book

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/chaihaobo/go-graphql-template/store"
)

type (
	Resolver interface {
		ID() graphql.ID
		Title() string
		Author() string
	}
	resolver struct {
		book *store.Book
	}
)

func (r resolver) ID() graphql.ID {
	return graphql.ID(r.book.ID)
}

func (r resolver) Title() string {
	return r.book.Title
}

func (r resolver) Author() string {
	return r.book.Author
}

func NewResolver(book *store.Book) Resolver {
	return &resolver{book: book}
}
