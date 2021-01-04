package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	authenticationService "example/api/grpc/authentication"
	"example/services/graphql/graph/graph_model"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func (r *mutationResolver) Register(ctx context.Context, input graph_model.Register) (*graph_model.User, error) {
	connection, err := grpc.Dial(os.Getenv("AUTHENTICATION_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to authentication server: %v", err)
		return &graph_model.User{}, err
	}
	defer connection.Close()
	client := authenticationService.NewAuthenticationServiceClient(connection)
	result, err := client.Register(ctx, &authenticationService.RegisterRequest{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.User{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	return &graph_model.User{
		Token:      &result.Token,
		ID:         &id,
		FirstName:  &result.FirstName,
		LastName:   &result.LastName,
		Email:      &result.Email,
		FacebookID: &result.FacebookId,
		GoogleID:   &result.GoogleId,
		Avatar:     &result.Avatar,
		CreatedAt:  &result.CreatedAt,
		UpdatedAt:  &result.UpdatedAt,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input graph_model.Login) (*graph_model.User, error) {
	connection, err := grpc.Dial(os.Getenv("AUTHENTICATION_GRPC_SERVER"), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Did not connect to authentication server: %v", err)
		return &graph_model.User{}, err
	}
	defer connection.Close()
	client := authenticationService.NewAuthenticationServiceClient(connection)
	result, err := client.Login(ctx, &authenticationService.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		log.Error(err)
		errorDetail, _ := status.FromError(err)
		return &graph_model.User{}, fmt.Errorf("%s", errorDetail.Message())
	}
	id := int(result.Id)
	return &graph_model.User{
		Token:      &result.Token,
		ID:         &id,
		FirstName:  &result.FirstName,
		LastName:   &result.LastName,
		Email:      &result.Email,
		FacebookID: &result.FacebookId,
		GoogleID:   &result.GoogleId,
		Avatar:     &result.Avatar,
		CreatedAt:  &result.CreatedAt,
		UpdatedAt:  &result.UpdatedAt,
	}, nil
}
