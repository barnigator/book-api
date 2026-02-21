package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string `env:"PORT"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("can't load .env")
	}

	portNum := os.Getenv("PORT")

	port := fmt.Sprintf(":%s", portNum)

	return &Config{port}
}
