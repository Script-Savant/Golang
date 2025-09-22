package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ConsumerKey string
	ConsumerSecret string
}

func InitConfig() *Config{
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env variables")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	return &Config{
		ConsumerKey: consumerKey,
		ConsumerSecret: consumerSecret,
	}

}