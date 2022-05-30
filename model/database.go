package model

import (
	"github.com/r0x16/PowerstoreOVMRestore/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	var err error
	db, err = gorm.Open(sqlite.Open(config.Global.DB), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	conn, err := db.DB()
	if err != nil {
		return err
	}
	return conn.Close()
}
