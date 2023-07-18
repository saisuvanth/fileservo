package app

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	environment := os.Getenv("APP_ENV")

	if environment == "dev" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
