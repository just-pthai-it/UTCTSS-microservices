package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Connection struct {
	Db     *sql.DB
	GormDb *gorm.DB
}

func (connection *Connection) Connect() {
	//var cfg mysql.Config
	//var dataSourceString string
	//var databaseType string
	//if os.Getenv("IS_DOCKERIZE") == "" {
	//	cfg = mysql.Config{
	//		User:                 "user",
	//		Passwd:               "password",
	//		Net:                  "tcp",
	//		Addr:                 "localhost" + ":" + "2101",
	//		DBName:               "nckh",
	//		AllowNativePasswords: true,
	//		ParseTime:            true,
	//	}
	//} else {
	//	if os.Getenv("DB_TYPE") == "mysql" || os.Getenv("DB_TYPE") == "mariadb" {
	//		cfg = mysql.Config{
	//			User:                 os.Getenv("DB_USERNAME"),
	//			Passwd:               os.Getenv("DB_PASSWORD"),
	//			Net:                  "tcp",
	//			Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
	//			DBName:               os.Getenv("DB_NAME"),
	//			AllowNativePasswords: true,
	//		}
	//
	//		dataSourceString = cfg.FormatDSN()
	//		databaseType = "mysql"
	//	}
	//
	//	if os.Getenv("DB_TYPE") == "postgres" {
	//		dataSourceString = "postgresql://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	//	}
	//}
	//
	//// Get a database handle.
	//var err error
	//connection.Db, err = sql.Open(databaseType, dataSourceString)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	for {
		var err error
		var dsn string
		if os.Getenv("IS_DOCKERIZE") == "" {
			dsn = "host=localhost user=postgres password=password dbname=ESS port=5432 sslmode=disable"
		} else {
			dsn = "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
		}

		connection.GormDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
		if err == nil {
			break
		}

		fmt.Println(err)
		time.Sleep(time.Second * 10)
	}
	//connection.GormDb, err = gorm.Open(gormMysql.New(gormMysql.Config{
	//	Conn: connection.Db,
	//}), &gorm.Config{})
}
