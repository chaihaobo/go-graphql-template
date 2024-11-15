package query

import (
	"context"

	"github.com/chaihaobo/go-graphql-template/resolver/query/user"
	usertype "github.com/chaihaobo/go-graphql-template/resolver/types/user"
	"github.com/chaihaobo/go-graphql-template/store"
)

type (
	QResolver interface {
		Users(ctx context.Context) []user.Resolver
		User(ctx context.Context, args *usertype.GetUserArgs) user.Resolver
	}
	resolver struct {
	}
)

func (r resolver) Users(ctx context.Context) []user.Resolver {
	result := make([]user.Resolver, 0)
	for _, userEntity := range store.Users() {
		result = append(result, user.NewResolver(userEntity))
	}
	return result
}

func (r resolver) User(ctx context.Context, args *usertype.GetUserArgs) user.Resolver {
	id := args.ID
	return user.NewResolver(store.GetUser(string(id)))
}

func NewResolver() QResolver {
	return &resolver{}
}
