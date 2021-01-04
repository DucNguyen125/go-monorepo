package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/services/graphql/graph/graph_model"
	"fmt"
)

func (r *mutationResolver) Register(ctx context.Context, input graph_model.Register) (*graph_model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input graph_model.Login) (*graph_model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
