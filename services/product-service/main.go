package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"example/services/product-service/services"
	"example/services/product-service/utils/mysql_util"
	"example/utils/logger"

	pb "example/api/grpc/product"

	"example/services/product-service/middlewares"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
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
	port := os.Getenv("PRODUCT_SERVER_PORT")
	if port == "" {
		port = "3002"
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
	pb.RegisterProductServiceServer(s, &services.ProductServer{})
	log.Printf("[INFO] start http server listening %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
