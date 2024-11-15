package user

import (
	"github.com/graph-gophers/graphql-go"

	"github.com/chaihaobo/go-graphql-template/resolver/query/book"
	"github.com/chaihaobo/go-graphql-template/store"
)

type (
	Resolver interface {
		ID() graphql.ID
		Name() string
		Books() []book.Resolver
	}
	resolver struct {
		user *store.User
	}
)

func (r *resolver) Books() []book.Resolver {
	result := make([]book.Resolver, 0)
	for _, bookEntity := range store.Books() {
		if bookEntity.UserID == string(r.user.ID) {
			result = append(result, book.NewResolver(bookEntity))
		}
	}
	return result
}

func (r *resolver) ID() graphql.ID {
	return graphql.ID(r.user.ID)
}
func (r *resolver) Name() string {
	return r.user.Name
}

func NewResolver(user *store.User) Resolver {
	return &resolver{user: user}
}
