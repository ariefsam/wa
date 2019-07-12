package main

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var config Config

func init() {
	config = LoadConfig()

	DB(func(db *gorm.DB) {
		db.AutoMigrate(&WhatsappMessage{})
	})

}

func DB(a func(*gorm.DB)) {
	db, err := gorm.Open("mysql", config.DB.Username+":"+config.DB.Password+"@/"+config.DB.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	a(db)
	db.Close()
	return
}
