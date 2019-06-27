package main

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

var config Config

func init() {
	config = LoadConfig()

	db, err := gorm.Open("mysql", "root:password@/whatsapp_cumi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect databasea")
	}
	defer db.Close()

	db.AutoMigrate(&WhatsappMessage{})
}

func DB(a func(*gorm.DB)) {
	db, err := gorm.Open("mysql", "root:password@/whatsapp_cumi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	a(db)
	return
}
