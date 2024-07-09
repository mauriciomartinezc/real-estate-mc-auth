package config

import "github.com/joho/godotenv"

func LoadEnv() error {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		return err
	}

	return nil
}
