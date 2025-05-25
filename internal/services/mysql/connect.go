package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	MAX_IDLE_DB_CONNECTION = 10
	MAX_OPEN_DB_CONNECTION = 100
)

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Connect without specifying database
	dsnInit := fmt.Sprintf("%s:%s@tcp(%s)/?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost)
	sqlDB, err := sql.Open("mysql", dsnInit)
	if err != nil {
		log.Fatal("Initial DB connection failed:", err)
	}

	// Create the database if it doesn't exist
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}
	sqlDB.Close()

	// Now connect using gorm with the database specified
	dsnFull := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsnFull), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM connection failed:", err)
	}

	sqlDB2, err := db.DB()
	if err != nil {
		log.Fatal("sql.DB from gorm.DB failed:", err)
	}
	sqlDB2.SetMaxIdleConns(MAX_IDLE_DB_CONNECTION)
	sqlDB2.SetMaxOpenConns(MAX_OPEN_DB_CONNECTION)
	sqlDB2.SetConnMaxLifetime(time.Hour)

	log.Printf("DB CONNECTED: %s", dbHost)

	DB = db
}
