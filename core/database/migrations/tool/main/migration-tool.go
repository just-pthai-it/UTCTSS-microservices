package main

import (
	"UTCTSS-microservices/core/database/migrations/tool"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	args := os.Args
	command := args[2]
	migrationsDirRelativePath := args[1]
	tool.SynchronizeDatabase(command, migrationsDirRelativePath+os.Getenv("SERVICE_NAME"))
}
