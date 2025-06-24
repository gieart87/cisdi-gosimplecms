package configs

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gosimplecms/utils/env"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var err error

	dbHost := env.GetEnv("DB_HOST", "localhost")
	dbPort := env.GetEnv("DB_PORT", "3306")
	dbUser := env.GetEnv("DB_USER", "root")
	dbPass := env.GetEnv("DB_PASS", "root")
	dbName := env.GetEnv("DB_NAME", "gosimplecms")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	DB = database

	return DB
}
