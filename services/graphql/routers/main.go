package routers

import (
	"example/services/graphql/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// router.Use(gin.Logger())
	router.Use(middlewares.Logger)
	router.Use(gin.Recovery())

	router.POST("/graphql", GraphqlHandler())
	if os.Getenv("RUN_MODE") != "release" {
		router.GET("/graphql", PlaygroundHandler())
	}

	return router
}
