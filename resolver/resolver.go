package resolver

import (
	"github.com/chaihaobo/go-graphql-template/resolver/mutation"
	"github.com/chaihaobo/go-graphql-template/resolver/query"
)

type (
	Resolver interface {
		query.QResolver
		mutation.MResolver
	}
	resolver struct {
		query.QResolver
		mutation.MResolver
	}
)

func New() Resolver {
	return &resolver{
		query.NewResolver(),
		mutation.NewResolver(),
	}
}
