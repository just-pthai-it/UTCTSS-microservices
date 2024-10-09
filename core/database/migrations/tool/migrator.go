package tool

import (
	"context"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"time"
)

func MigrateUpDatabase(migrationsDirRelativePath string) error {
	context_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	dataSourceString := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := goose.OpenDBWithDriver(os.Getenv("DB_TYPE"), dataSourceString)
	if err != nil {
		log.Printf("failed to open database: %v\n", err)
		return err
	}

	if err := goose.UpContext(context_, db, migrationsDirRelativePath); err != nil {
		log.Printf("failed to up migrations: %v", err)
		return err
	}

	return nil
}

func SynchronizeDatabase(command string, migrationsDirRelativePath string) {
	context_, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	dataSourceString := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	db, err := goose.OpenDBWithDriver(os.Getenv("DB_TYPE"), dataSourceString)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := goose.RunContext(context_, command, db, migrationsDirRelativePath); err != nil {
		log.Fatalf("failed to %s migrations: %v", command, err)
	}
}
