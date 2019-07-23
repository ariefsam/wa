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
		Driver     string
		Connection string
	}
}

func LoadConfig() (data Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.ReceivePushURL = os.Getenv("RECEIVE_PUSH_URL")
	data.SendPopURL = os.Getenv("SEND_POP_URL")
	data.DB.Driver = os.Getenv("DB_DRIVER")
	data.DB.Connection = os.Getenv("DB_CONNECTION")
	return data
}
