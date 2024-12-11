package subscription

import (
	"context"

	"github.com/chaihaobo/go-graphql-template/resolver/query/user"
)

type (
	SResolver interface {
		UserCreated(ctx context.Context) <-chan user.Resolver
	}
	resolver struct {
		userCreatedChan <-chan user.Resolver
	}
)

func (r resolver) UserCreated(ctx context.Context) <-chan user.Resolver {
	return r.userCreatedChan
}

func NewResolver(userCreatedChan <-chan user.Resolver) SResolver {
	return resolver{userCreatedChan: userCreatedChan}
}
