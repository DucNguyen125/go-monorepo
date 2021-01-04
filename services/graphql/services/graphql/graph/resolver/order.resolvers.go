package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/services/graphql/graph/custom_model"
	"example/services/graphql/graph/graph_model"
	"fmt"
)

func (r *mutationResolver) OrderCreate(ctx context.Context, input graph_model.CreateOrder) (*graph_model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderUpdate(ctx context.Context, input custom_model.UpdateOrder) (*graph_model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) OrderDelete(ctx context.Context, input graph_model.DeleteOrder) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id int) (*graph_model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context, filter string, limit int, page int) ([]graph_model.Order, error) {
	panic(fmt.Errorf("not implemented"))
}
