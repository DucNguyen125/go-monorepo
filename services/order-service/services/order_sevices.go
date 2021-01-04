package services

import (
	"context"
	pb "example/api/grpc/order"
	"example/services/order-service/models"
	"example/services/order-service/utils/mysql_util"
	"time"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
}

func (server *OrderServer) CreateOrder(context context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	newOrder := models.Order{
		OrderCode:   request.GetOrderCode(),
		OrderType:   request.GetOrderType(),
		Products:    request.GetProducts(),
		OrderStatus: request.GetOrderStatus(),
		Quantity:    int(request.GetQuantity()),
		TotalPrice:  int(request.TotalPrice),
	}
	if err := mysql_util.DB.Create(&newOrder).Error; err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	createdOrder := &pb.CreateOrderResponse{
		Id:          int32(newOrder.ID),
		OrderCode:   newOrder.OrderCode,
		OrderType:   newOrder.OrderType,
		Products:    newOrder.Products,
		OrderStatus: newOrder.OrderStatus,
		Quantity:    int32(newOrder.Quantity),
		TotalPrice:  int64(newOrder.TotalPrice),
		CreatedAt:   newOrder.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newOrder.UpdatedAt.Format(time.RFC3339),
	}
	return createdOrder, nil
}

func (server *OrderServer) UpdateOrder(context context.Context, request *pb.UpdateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &pb.CreateOrderResponse{}
	err := mysql_util.DB.Updates(models.Order{
		ID:          int(request.GetId()),
		OrderCode:   request.GetOrderCode(),
		OrderType:   request.GetOrderType(),
		Products:    request.GetProducts(),
		OrderStatus: request.GetOrderStatus(),
		Quantity:    int(request.GetQuantity()),
		TotalPrice:  int(request.TotalPrice),
	}).First(order, int(request.GetId())).Error
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return order, nil
}

func (server *OrderServer) GetOrder(context context.Context, request *pb.GetOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &pb.CreateOrderResponse{}
	err := mysql_util.DB.Model(&models.Order{}).First(&order, request.GetId()).Error
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}
	return order, nil
}

func (server *OrderServer) GetListOrder(context context.Context, request *pb.GetListOrderRequest) (*pb.GetListOrderResponse, error) {
	var listOrder []*pb.CreateOrderResponse
	limit := int(request.GetLimit())
	skip := int(request.GetPage()-1) * limit
	mysql_util.DB.Model(&models.Order{}).Offset(skip).Limit(limit).Find(&listOrder)
	return &pb.GetListOrderResponse{ListOrder: listOrder}, nil
}

func (server *OrderServer) DeleteOrder(context context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	err := mysql_util.DB.Delete(&models.Order{}, request.GetId()).Error
	if err != nil {
		return &pb.DeleteOrderResponse{Message: "Delete order failed"}, err
	} else {
		return &pb.DeleteOrderResponse{Message: "Delete order success"}, nil
	}
}
