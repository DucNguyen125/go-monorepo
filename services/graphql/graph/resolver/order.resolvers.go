package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	orderService "example/api/grpc/order"
	"example/services/graphql/graph/custom_model"
	"example/services/graphql/graph/graph_model"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (r *mutationResolver) OrderCreate(ctx context.Context, input graph_model.CreateOrder) (*graph_model.Order, error) {
	connection, err := grpc.Dial(os.Getenv("ORDER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to order server: %v", err)
		return &graph_model.Order{}, err
	}
	defer connection.Close()
	client := orderService.NewOrderServiceClient(connection)
	result, err := client.CreateOrder(ctx, &orderService.CreateOrderRequest{
		OrderCode:   input.OrderCode,
		OrderType:   input.OrderType,
		Products:    input.Products,
		OrderStatus: input.OrderStatus,
		Quantity:    int32(input.Quantity),
		TotalPrice:  int64(input.TotalPrice),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Order{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	quantity := int(result.Quantity)
	totalPrice := int(result.TotalPrice)
	return &graph_model.Order{
		ID:          &id,
		OrderCode:   &result.OrderCode,
		OrderType:   &result.OrderType,
		Products:    &result.Products,
		OrderStatus: &result.OrderStatus,
		Quantity:    &quantity,
		TotalPrice:  &totalPrice,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *mutationResolver) OrderUpdate(ctx context.Context, input custom_model.UpdateOrder) (*graph_model.Order, error) {
	connection, err := grpc.Dial(os.Getenv("ORDER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to order server: %v", err)
		return &graph_model.Order{}, err
	}
	defer connection.Close()
	client := orderService.NewOrderServiceClient(connection)
	result, err := client.UpdateOrder(ctx, &orderService.UpdateOrderRequest{
		Id:          int32(input.ID),
		OrderCode:   input.OrderCode,
		OrderType:   input.OrderType,
		Products:    input.Products,
		OrderStatus: input.OrderStatus,
		Quantity:    int32(input.Quantity),
		TotalPrice:  int64(input.TotalPrice),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Order{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	quantity := int(result.Quantity)
	totalPrice := int(result.TotalPrice)
	return &graph_model.Order{
		ID:          &id,
		OrderCode:   &result.OrderCode,
		OrderType:   &result.OrderType,
		Products:    &result.Products,
		OrderStatus: &result.OrderStatus,
		Quantity:    &quantity,
		TotalPrice:  &totalPrice,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *mutationResolver) OrderDelete(ctx context.Context, input graph_model.DeleteOrder) (*int, error) {
	connection, err := grpc.Dial(os.Getenv("ORDER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to order server: %v", err)
		return nil, err
	}
	defer connection.Close()
	client := orderService.NewOrderServiceClient(connection)
	_, err = client.DeleteOrder(ctx, &orderService.DeleteOrderRequest{
		Id: int32(input.ID),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return nil, fmt.Errorf("%s", errorDetail.Message())
	}
	return &input.ID, nil
}

func (r *queryResolver) Order(ctx context.Context, id int) (*graph_model.Order, error) {
	connection, err := grpc.Dial(os.Getenv("ORDER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to order server: %v", err)
		return &graph_model.Order{}, err
	}
	defer connection.Close()
	client := orderService.NewOrderServiceClient(connection)
	result, err := client.GetOrder(ctx, &orderService.GetOrderRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.Order{}, fmt.Errorf("%s", errorDetail.Message())
	}
	orderId := int(result.Id)
	quantity := int(result.Quantity)
	totalPrice := int(result.TotalPrice)
	return &graph_model.Order{
		ID:          &orderId,
		OrderCode:   &result.OrderCode,
		OrderType:   &result.OrderType,
		Products:    &result.Products,
		OrderStatus: &result.OrderStatus,
		Quantity:    &quantity,
		TotalPrice:  &totalPrice,
		CreatedAt:   &result.CreatedAt,
		UpdatedAt:   &result.UpdatedAt,
	}, nil
}

func (r *queryResolver) Orders(ctx context.Context, filter string, limit int, page int) ([]graph_model.Order, error) {
	connection, err := grpc.Dial(os.Getenv("ORDER_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to order server: %v", err)
		return []graph_model.Order{}, err
	}
	defer connection.Close()
	client := orderService.NewOrderServiceClient(connection)
	result, err := client.GetListOrder(ctx, &orderService.GetListOrderRequest{
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return []graph_model.Order{}, fmt.Errorf("%s", errorDetail.Message())
	}
	listOrder := []graph_model.Order{}
	for _, order := range result.ListOrder {
		id := int(order.Id)
		quantity := int(order.Quantity)
		totalPrice := int(order.TotalPrice)
		listOrder = append(listOrder, graph_model.Order{
			ID:          &id,
			OrderCode:   &order.OrderCode,
			OrderType:   &order.OrderType,
			Products:    &order.Products,
			OrderStatus: &order.OrderStatus,
			Quantity:    &quantity,
			TotalPrice:  &totalPrice,
			CreatedAt:   &order.CreatedAt,
			UpdatedAt:   &order.UpdatedAt,
		})
	}
	return listOrder, nil
}
