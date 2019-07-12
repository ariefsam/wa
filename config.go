package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ReceivePushURL string
	SendPopURL     string
	DB             struct {
		Database string
		Username string
		Password string
	}
}

func LoadConfig() (data Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.ReceivePushURL = os.Getenv("RECEIVE_PUSH_URL")
	data.SendPopURL = os.Getenv("SEND_POP_URL")
	data.DB.Username = os.Getenv("DB_USERNAME")
	data.DB.Password = os.Getenv("DB_PASSWORD")
	data.DB.Database = os.Getenv("DB_DATABASE")
	return data
}
