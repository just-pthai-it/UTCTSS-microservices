package main

import (
	"ESS-microservices/core/middlewares"
	"ESS-microservices/core/routes"
	"ESS-microservices/database/migrations/tool"
	_ "ESS-microservices/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	tool.MigrateDatabase("database/migrations/" + os.Getenv("SERVICE_NAME"))

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
