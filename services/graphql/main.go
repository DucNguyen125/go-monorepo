package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example/services/graphql/routers"
	"example/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("RUN_MODE"))

	// Init logger
	logger.Init()

	port := os.Getenv("GRAPHQL_SERVER_PORT")
	if port == "" {
		port = "3000"
	}
	// Register Router
	routersInit := routers.InitRouter()
	if err != nil {
		log.Fatal("Error get PORT")
	}
	endPoint := fmt.Sprintf(":%s", port)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err = server.ListenAndServe(); err != nil {
		log.Fatal("Fail to start error server")
	}
}
