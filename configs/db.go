package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gosimplecms/utils/env"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error

	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "3306")
	dbUser := env.GetEnv("DB_USER", "root")
	dbPass := env.GetEnv("DB_PASS", "root")
	dbName := env.GetEnv("DB_NAME", "gosimplecms")

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)
	}

	for i := 0; i < 10; i++ {
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = database
			fmt.Println("Connected to DB!")
			return DB
		}
		log.Printf("Retry DB connection... (%d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Failed to connect database after retries: %v", err)

	return DB
}
