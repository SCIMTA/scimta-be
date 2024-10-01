package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"localhost", "postgres", "123456", "scimta_be", "5432", "disable", "Asia/Tokyo")
}

func New() *gorm.DB {
	// Connect to database
	dsn := GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger.Info(context.Background(), "Connected to database")

	if err != nil {
		log.Panic(err)
		panic(err)
	}

	dbConfig, err := db.DB()
	dbConfig.SetMaxIdleConns(10)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	db.Logger.LogMode(logger.Info)

	return db
}

func TestDB() *gorm.DB {
	// Connect to database
	dsn := GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	dbConfig, err := db.DB()
	dbConfig.SetMaxIdleConns(10)
	if err != nil {
		log.Panic(err)
		panic(err)
	}

	db.Logger.LogMode(logger.Info)
	return db
}

func DropTestDB() error {
	db, err := gorm.Open(postgres.Open(GetPostgresDSN()))
	if err != nil {
		log.Panic(err)
	}
	db.Exec("DROP DATABASE IF EXIST scimta_be")
	return nil
}
