package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTIONSTRING")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ")
	}

	DB.AutoMigrate(&User{})

}
