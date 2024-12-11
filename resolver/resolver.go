package resolver

import (
	"github.com/chaihaobo/go-graphql-template/resolver/mutation"
	"github.com/chaihaobo/go-graphql-template/resolver/query"
	"github.com/chaihaobo/go-graphql-template/resolver/query/user"
	"github.com/chaihaobo/go-graphql-template/resolver/subscription"
)

type (
	Resolver interface {
		query.QResolver
		mutation.MResolver
		subscription.SResolver
	}
	resolver struct {
		query.QResolver
		mutation.MResolver
		subscription.SResolver
	}
)

func New() Resolver {
	userCreatedChan := make(chan user.Resolver)
	return &resolver{
		query.NewResolver(),
		mutation.NewResolver(userCreatedChan),
		subscription.NewResolver(userCreatedChan),
	}
}
