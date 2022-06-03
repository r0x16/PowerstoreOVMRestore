package database

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func Connect() error {
	var err error
	Connection, err = gorm.Open(sqlite.Open(config.Global.DB), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	conn, err := Connection.DB()
	if err != nil {
		return err
	}
	return conn.Close()
}
