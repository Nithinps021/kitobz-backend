package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func getEnvironmentVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error in loading .env file")
	}
	return os.Getenv(key)
}