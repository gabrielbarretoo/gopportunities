package config

import (
	"os"

	"github.com/gabrielbarretoo/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Database file does not exist, creating one")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}
	// Migrate schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to migrate schema: %v", err)
		return nil, err
	}

	// Return DB
	return db, nil
}
