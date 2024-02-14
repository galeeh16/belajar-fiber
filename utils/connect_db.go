package utils

import (
	"galih/belajar-fiber/models/entity"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Gagal koneksi ke database", err.Error())
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Contact{},
		&entity.Post{},
	)

	DB = db
}
