package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Connection struct {
	Db     *sql.DB
	GormDb *gorm.DB
}

func (connection *Connection) Connect() {
	fmt.Println()
	var cfg mysql.Config
	if os.Getenv("IS_DOCKERIZE") == "" {
		cfg = mysql.Config{
			User:                 "user",
			Passwd:               "password",
			Net:                  "tcp",
			Addr:                 "localhost" + ":" + "2101",
			DBName:               "nckh",
			AllowNativePasswords: true,
			ParseTime:            true,
		}
	} else {
		cfg = mysql.Config{
			User:                 os.Getenv("DB_USERNAME"),
			Passwd:               os.Getenv("DB_PASSWORD"),
			Net:                  "tcp",
			Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
			DBName:               os.Getenv("DB_NAME"),
			AllowNativePasswords: true,
		}
	}

	// Get a database handle.
	var err error
	connection.Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
		return
	}

	connection.GormDb, err = gorm.Open(gormMysql.New(gormMysql.Config{
		Conn: connection.Db,
	}), &gorm.Config{})
}
