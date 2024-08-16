package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func New(uri string, isDevelopment bool) (*gorm.DB, error) {
	if db != nil {
		log.Println("Database already initialized, returning the same instance")
		return db, nil
	}

	client, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to client", "error", err)
		return nil, err
	}

	configSQLDriver, err := client.DB()
	if err != nil {
		log.Fatal("Error getting SQL driver", "error", err)
		return nil, err
	}

	configSQLDriver.SetMaxIdleConns(10)
	configSQLDriver.SetMaxOpenConns(50)
	configSQLDriver.SetConnMaxIdleTime(30 * time.Minute)
	configSQLDriver.SetConnMaxLifetime(time.Hour)

	if isDevelopment {
		client = client.Debug()
	}

	db = client

	return client, nil
}

func Disconnect() {
	if db == nil {
		log.Println("Database not initialized")
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting SQL driver", "error", err)
	}

	sqlDB.Close()
}

func GetDB() *gorm.DB {
	if db == nil {
		panic("Database not initialized")
	}

	return db
}

func Migrate(dst ...interface{}) error {
	if db == nil {
		panic("Database not initialized")
	}

	err := db.AutoMigrate(dst...)
	if err != nil {
		log.Fatal("Error migrating database", "error", err)
		return err
	}

	return nil
}
