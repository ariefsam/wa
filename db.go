package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var config Config

func init() {
	config = LoadConfig()

	DB(func(db *gorm.DB) {
		db.AutoMigrate(&WhatsappMessage{})
	})

}

func DB(a func(*gorm.DB)) {
	db, err := gorm.Open(config.DB.Driver, config.DB.Connection)
	if err != nil {
		fmt.Println(config.DB.Driver, config.DB.Connection)
		panic("failed to connect database")
	}
	a(db)
	db.Close()
	return
}
