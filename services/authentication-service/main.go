package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"example/services/authentication-service/services"
	"example/services/authentication-service/utils/mysql_util"
	"example/utils/logger"

	pb "example/api/grpc/authentication"

	"example/services/authentication-service/middlewares"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}

	// Init logger
	logger.Init()

	// Connect to database
	if err = mysql_util.Connect(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Migrate database
	if err = mysql_util.AutoMigrate(); err != nil {
		log.Fatal("Error migrating to database")
	}
	port := os.Getenv("AUTHENTICATION_SERVER_PORT")
	if port == "" {
		port = "3003"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			middlewares.Logger(),
		)),
	)
	pb.RegisterAuthenticationServiceServer(s, &services.AuthenticationServer{})
	log.Printf("[INFO] start http server listening %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
