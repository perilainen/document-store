package models

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("documents.db"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	db.AutoMigrate(&Document{})
	DB = db

}
