package services

import (
	"context"
	pb "example/api/grpc/authentication"
	"example/services/authentication-service/models"
	"example/services/authentication-service/utils/encrypt"
	jwt_service "example/services/authentication-service/utils/jwt"
	"example/services/authentication-service/utils/mysql_util"
	"time"
)

type AuthenticationServer struct {
	pb.UnimplementedAuthenticationServiceServer
}

func (server *AuthenticationServer) Register(context context.Context, request *pb.RegisterRequest) (*pb.UserResponse, error) {
	hashPassword, err := encrypt.HashPassword(request.GetPassword())
	if err != nil {
		return &pb.UserResponse{}, err
	}
	newUser := models.User{
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
		Email:     request.GetEmail(),
		Password:  hashPassword,
	}
	if err := mysql_util.DB.Create(&newUser).Error; err != nil {
		return &pb.UserResponse{}, err
	}
	token := jwt_service.GenerateToken(newUser.ID)
	createdUser := &pb.UserResponse{
		Token:      token,
		Id:         int32(newUser.ID),
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		Email:      newUser.Email,
		FacebookId: newUser.FacebookId,
		GoogleId:   newUser.GoogleId,
		Avatar:     newUser.Avatar,
		CreatedAt:  newUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  newUser.UpdatedAt.Format(time.RFC3339),
	}
	return createdUser, nil
}

func (server *AuthenticationServer) Login(context context.Context, request *pb.LoginRequest) (*pb.UserResponse, error) {
	user := models.User{}
	err := mysql_util.DB.Model(&models.User{}).Where("email = ?", request.GetEmail()).First(&user).Error
	if err != nil {
		return &pb.UserResponse{}, err
	}
	if err = encrypt.CheckPassword(request.GetPassword(), user.Password); err != nil {
		return &pb.UserResponse{}, err
	}
	token := jwt_service.GenerateToken(user.ID)
	result := &pb.UserResponse{
		Token:      token,
		Id:         int32(user.ID),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		FacebookId: user.FacebookId,
		GoogleId:   user.GoogleId,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}
	return result, nil
}
