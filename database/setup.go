package database

import (
	"fmt"
	"log"
	"os"
	"todoapp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// gorm is the orm for go
var DB *gorm.DB

func SetupDatabase() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		"localhost", // host
		"postgres",  // user
		"",  // password
		"todoapp",   // dbname
		"5432",      // port
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		os.Exit(1)
	}
}