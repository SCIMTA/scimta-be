package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"scimta-be/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GetPostgresDSN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PORT"),
		"disable",
		"Asia/Tokyo")
}

func New() *gorm.DB {
	// Connect to database
	dsn := GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "public.", // schema name
			SingularTable: false,
		}})
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

func AutoMigrate(db *gorm.DB) {
	log.Default().Println("Start migrating database")
	db.AutoMigrate(&model.User{})
}

func DropTestDB() error {
	db, err := gorm.Open(postgres.Open(GetPostgresDSN()))
	if err != nil {
		log.Panic(err)
	}
	db.Exec("DROP DATABASE IF EXIST scimta_be")
	return nil
}
