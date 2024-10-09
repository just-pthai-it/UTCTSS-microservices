package main

import (
	"UTCTSS-microservices/common"
	"UTCTSS-microservices/core/database/migrations/tool"
	"UTCTSS-microservices/core/http/middlewares"
	_ "UTCTSS-microservices/docs"
	"UTCTSS-microservices/services/teacher/http/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	migrateDatabase()

	router := gin.Default()
	router.Use(middlewares.ConsoleLoggerMiddleware)
	router.Use(middlewares.FileLoggerMiddleware)
	router.Use(middlewares.GinLoggerCustomFormat())

	routes.NewApiRoutes(router)

	// Run the server on port 8080
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func migrateDatabase() {
	databaseMigrateMaxAttempts := 10
	for {
		serviceName := os.Getenv("SERVICE_NAME")
		err := tool.MigrateUpDatabase("./services/" + serviceName + "/migrations")

		if err != nil {
			log.Println(common.Red + "Migrate failed: " + err.Error() + common.Reset)
		} else {
			log.Println(common.Green + "Migrate successful!" + common.Reset)
			break
		}

		databaseMigrateMaxAttempts--
		if databaseMigrateMaxAttempts == 0 {
			log.Fatal(common.Red + "Exit because of database migrate failed!" + common.Reset)
		}

		time.Sleep(time.Second)
	}
}
