package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lujakob/gift-sats/tip"
	"github.com/lujakob/gift-sats/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(dsn string) *gorm.DB {
	//dsn := "host=/tmp user=realworld dbname=realworld"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	/*
	 *db, err := gorm.Open(postgres.New(postgres.Config{
	 *  DSN: dsn,
	 *  //PreferSimpleProtocol: true, // disables implicit prepared statement usage
	 *}), &gorm.Config{})
	 */

	//db, err := gorm.Open("postgresql", "postgresql://realworld@/realworld?host=/tmp")
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
		&user.User{},
		&tip.Tip{},
	)
}
