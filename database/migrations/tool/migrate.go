package tool

import (
	_ "TSS-microservices/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func MigrateDatabase(migrationsFilepath string) {
	dataSourceString := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := goose.OpenDBWithDriver(os.Getenv("DB_TYPE"), dataSourceString)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	fmt.Println(dataSourceString)
	if err := goose.Up(db, migrationsFilepath); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
