package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func DbConnect() {
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Database not found")
	}

	err = db.AutoMigrate(&User{})

	if err != nil {
		log.Fatal("nao foi possivel migrar")
	}

	Db = db
}
