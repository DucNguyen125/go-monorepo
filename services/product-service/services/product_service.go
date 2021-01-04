package services

import (
	"context"
	pb "example/api/grpc/product"
	"example/services/product-service/models"
	"example/services/product-service/utils/mysql_util"
	"time"
)

type ProductServer struct {
	pb.UnimplementedProductServiceServer
}

func (server *ProductServer) CreateProduct(context context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	newProduct := models.Product{
		ProductCode: request.GetProductCode(),
		ProductName: request.GetProductName(),
		Price:       int(request.GetPrice()),
	}
	if err := mysql_util.DB.Create(&newProduct).Error; err != nil {
		return &pb.CreateProductResponse{}, err
	}
	createdProduct := &pb.CreateProductResponse{
		Id:          int32(newProduct.ID),
		ProductCode: newProduct.ProductCode,
		ProductName: newProduct.ProductName,
		Price:       int64(newProduct.Price),
		CreatedAt:   newProduct.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   newProduct.UpdatedAt.Format(time.RFC3339),
	}
	return createdProduct, nil
}

func (server *ProductServer) UpdateProduct(context context.Context, request *pb.UpdateProductRequest) (*pb.CreateProductResponse, error) {
	product := &pb.CreateProductResponse{}
	err := mysql_util.DB.Updates(models.Product{
		ID:          int(request.GetId()),
		ProductCode: request.GetProductCode(),
		ProductName: request.GetProductName(),
		Price:       int(request.GetPrice()),
	}).First(product, int(request.GetId())).Error
	if err != nil {
		return &pb.CreateProductResponse{}, err
	}
	return product, nil
}

func (server *ProductServer) GetProduct(context context.Context, request *pb.GetProductRequest) (*pb.CreateProductResponse, error) {
	product := &pb.CreateProductResponse{}
	err := mysql_util.DB.Model(&models.Product{}).First(&product, request.GetId()).Error
	if err != nil {
		return &pb.CreateProductResponse{}, err
	}
	return product, nil
}

func (server *ProductServer) GetListProduct(context context.Context, request *pb.GetListProductRequest) (*pb.GetListProductResponse, error) {
	var listProduct []*pb.CreateProductResponse
	limit := int(request.GetLimit())
	skip := int(request.GetPage()-1) * limit
	mysql_util.DB.Model(&models.Product{}).Offset(skip).Limit(limit).Find(&listProduct)
	return &pb.GetListProductResponse{ListProduct: listProduct}, nil
}

func (server *ProductServer) DeleteProduct(context context.Context, request *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	err := mysql_util.DB.Delete(&models.Product{}, request.GetId()).Error
	if err != nil {
		return &pb.DeleteProductResponse{Message: "Delete product failed"}, err
	} else {
		return &pb.DeleteProductResponse{Message: "Delete product success"}, nil
	}
}
