package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./config/env_vars.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println(err.Error())
	}
}
