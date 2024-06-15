package main

import (
	"ESS-microservices/core/middlewares"
	"ESS-microservices/core/routes"
	_ "ESS-microservices/docs"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(middlewares.ConsoleLoggerMiddleware)
	router.Use(middlewares.FileLoggerMiddleware)

	routes.NewApiRoutes(router)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
