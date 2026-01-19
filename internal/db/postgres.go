package db

import (	
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectPostgres() {
	
   // its like a connection string in java and c#
	//dsn := "host=localhost user=postgres password=postgres dbname=health_monitor port=5432 sslmode=disable TimeZone=UTC"


   host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        host, user, password, dbname, port)

	//  grom.open (connection string , conifig ) returns two values (db object and error)
		database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	DB = database

	log.Println("Postgres Connected Successfully")
}