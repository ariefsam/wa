package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	ReceivePushURL string
	SendPopURL     string
}

func LoadConfig() (data Config) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	data.Port = os.Getenv("PORT")
	data.ReceivePushURL = os.Getenv("RECEIVE_PUSH_URL")
	data.SendPopURL = os.Getenv("SEND_POP_URL")
	return data
}
