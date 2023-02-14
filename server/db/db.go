package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lujakob/gift-sats/config"
	"github.com/lujakob/gift-sats/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New() *gorm.DB {
	config := config.GetConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	dsn := config.DB_DSN

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	//db, err := gorm.Open("sqlite3", "./database/realworld.db")
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func TestDB() *gorm.DB {
	dsn := "./../database/gift-sats_test.db"
	//newLogger := logger.New(
	//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//logger.Config{
	//SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
	//LogLevel:                  logger.Info,           // Log level
	//IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
	//Colorful:                  true,                  // Disable color
	//},
	//)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../database/gift-sats_test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Tip{},
		&models.Wallet{},
	)
}
