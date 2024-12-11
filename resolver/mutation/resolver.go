package mutation

import (
	"context"

	"github.com/chaihaobo/go-graphql-template/resolver/query/book"
	"github.com/chaihaobo/go-graphql-template/resolver/query/user"
	booktype "github.com/chaihaobo/go-graphql-template/resolver/types/book"
	usertype "github.com/chaihaobo/go-graphql-template/resolver/types/user"
	"github.com/chaihaobo/go-graphql-template/store"
)

type (
	MResolver interface {
		CreateUser(ctx context.Context, args *usertype.CreateUserArgs) (user.Resolver, error)
		CreateBook(ctx context.Context, args *booktype.CreateBookArgs) (book.Resolver, error)
	}
	resolver struct {
		userCreatedChan chan<- user.Resolver
	}
)

func (r resolver) CreateUser(ctx context.Context, args *usertype.CreateUserArgs) (user.Resolver, error) {
	userResolver := user.NewResolver(store.CreateUser(args.Input.Name))
	defer func() {
		go func() {
			r.userCreatedChan <- userResolver
		}()
	}()
	return userResolver, nil
}

func (r resolver) CreateBook(ctx context.Context, args *booktype.CreateBookArgs) (book.Resolver, error) {
	return book.NewResolver(store.CreateBook(args.Input.Author, args.Input.Title, args.Input.UserID)), nil
}

func NewResolver(userCreatedChan chan<- user.Resolver) MResolver {
	return resolver{userCreatedChan: userCreatedChan}
}
