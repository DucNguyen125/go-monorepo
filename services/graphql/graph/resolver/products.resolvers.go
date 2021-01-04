package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	productService "example/api/grpc/product"
	"example/services/graphql/graph/custom_model"
	"example/services/graphql/graph/graph_model"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (r *mutationResolver) ProductCreate(ctx context.Context, input graph_model.CreateProduct) (*graph_model.Product, error) {
	connection, err := grpc.Dial(os.Getenv("PRODUCT_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to product server: %v", err)
		return &graph_model.Product{}, err
	}
	defer connection.Close()
	client := productService.NewProductServiceClient(connection)
	result, err := client.CreateProduct(ctx, &productService.CreateProductRequest{
		ProductCode: input.ProductCode,
		ProductName: input.ProductName,
		Price:       int64(input.Price),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Product{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	price := int(result.Price)
	return &graph_model.Product{
		ID:          &id,
		ProductCode: &result.ProductCode,
		ProductName: &result.ProductName,
		Price:       &price,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *mutationResolver) ProductUpdate(ctx context.Context, input custom_model.UpdateProduct) (*graph_model.Product, error) {
	connection, err := grpc.Dial(os.Getenv("PRODUCT_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to product server: %v", err)
		return &graph_model.Product{}, err
	}
	defer connection.Close()
	client := productService.NewProductServiceClient(connection)
	result, err := client.UpdateProduct(ctx, &productService.UpdateProductRequest{
		Id:          int32(input.ID),
		ProductCode: input.ProductCode,
		ProductName: input.ProductName,
		Price:       int64(input.Price),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Product{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	price := int(result.Price)
	return &graph_model.Product{
		ID:          &id,
		ProductCode: &result.ProductCode,
		ProductName: &result.ProductName,
		Price:       &price,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *mutationResolver) ProductDelete(ctx context.Context, input graph_model.DeleteProduct) (*int, error) {
	connection, err := grpc.Dial(os.Getenv("PRODUCT_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to product server: %v", err)
		return nil, err
	}
	defer connection.Close()
	client := productService.NewProductServiceClient(connection)
	_, err = client.DeleteProduct(ctx, &productService.DeleteProductRequest{
		Id: int32(input.ID),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return nil, fmt.Errorf("%s", errorDetail.Message())
	}
	return &input.ID, nil
}

func (r *queryResolver) Product(ctx context.Context, id int) (*graph_model.Product, error) {
	connection, err := grpc.Dial(os.Getenv("PRODUCT_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to product server: %v", err)
		return &graph_model.Product{}, err
	}
	defer connection.Close()
	client := productService.NewProductServiceClient(connection)
	result, err := client.GetProduct(ctx, &productService.GetProductRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Product{}, fmt.Errorf("%s", errorDetail.Message())
	}
	productId := int(result.Id)
	price := int(result.Price)
	return &graph_model.Product{
		ID:          &productId,
		ProductCode: &result.ProductCode,
		ProductName: &result.ProductName,
		Price:       &price,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *queryResolver) Products(ctx context.Context, filter string, limit int, page int) ([]graph_model.Product, error) {
	connection, err := grpc.Dial(os.Getenv("PRODUCT_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to product server: %v", err)
		return []graph_model.Product{}, err
	}
	defer connection.Close()
	client := productService.NewProductServiceClient(connection)
	result, err := client.GetListProduct(ctx, &productService.GetListProductRequest{
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return []graph_model.Product{}, fmt.Errorf("%s", errorDetail.Message())
	}
	listProduct := []graph_model.Product{}
	for _, product := range result.ListProduct {
		id := int(product.Id)
		price := int(product.Price)
		listProduct = append(listProduct, graph_model.Product{
			ID:          &id,
			ProductCode: &product.ProductCode,
			ProductName: &product.ProductName,
			Price:       &price,
			CreatedAt:   &product.CreatedAt,
			UpdatedAt:   &product.UpdatedAt,
		})
	}
	return listProduct, nil
}
