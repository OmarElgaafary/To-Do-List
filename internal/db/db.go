package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var agorm *gorm.DB

func Init() error {
	var err error
	agorm, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return agorm
}
