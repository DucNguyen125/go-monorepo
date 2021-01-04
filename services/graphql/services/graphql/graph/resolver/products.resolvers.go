package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example/services/graphql/graph/custom_model"
	"example/services/graphql/graph/graph_model"
	"fmt"
)

func (r *mutationResolver) ProductCreate(ctx context.Context, input graph_model.CreateProduct) (*graph_model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, input custom_model.UpdateProduct) (*graph_model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ProductDelete(ctx context.Context, input graph_model.DeleteProduct) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id int) (*graph_model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, filter string, limit int, page int) ([]graph_model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}
