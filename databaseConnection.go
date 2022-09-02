package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() *gorm.DB {
	url := getEnvironmentVariable(DB_URL)
	db, err :=  gorm.Open(postgres.Open(url),&gorm.Config{})
	
	if err != nil{
		log.Fatal("Failed to connect to db ",err)
	}
	log.Fatal("Connected to DB ",db.Name())

	return db
}