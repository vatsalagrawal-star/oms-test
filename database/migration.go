package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"oms-test/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=user password=password dbname=omsdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Connected to database successfully!")	
}

func Migrate() {
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migration completed.")
}
