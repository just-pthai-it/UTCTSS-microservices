package main

import (
	"UTCTSS-microservices/common"
	"UTCTSS-microservices/core/middlewares"
	"UTCTSS-microservices/core/routes"
	_ "UTCTSS-microservices/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
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
		migrateCmd := exec.Command("./build/migrationtool", "./database/migrations/", "up")
		output, err := migrateCmd.CombinedOutput()

		if err != nil {
			fmt.Println(common.Red + "Migrate failed: " + err.Error() + common.Reset)
		} else {
			fmt.Println(common.Blue + "Migrate successful!" + common.Reset)
			fmt.Println(string(output))
			break
		}

		databaseMigrateMaxAttempts--
		if databaseMigrateMaxAttempts == 0 {
			if err != nil {
				log.Fatal(common.Red + "Exit because of database migrating failed!" + common.Reset)
			}
			break
		}
		time.Sleep(time.Second)
	}
}