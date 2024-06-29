package database

import (
	"TSS-microservices/database/drivers"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Connection struct {
	Driver drivers.IDriver
	Db     *sql.DB
	GormDb *gorm.DB
}

func (connection *Connection) Connect() {
	for {
		var err error
		connection.GormDb, err = connection.Driver.InitConnection()
		if err == nil {
			break
		}

		fmt.Println(err)
		time.Sleep(time.Second * 10)
	}
}
