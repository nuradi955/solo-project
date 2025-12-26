package data_base

import (
	"solo-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var dsn = "host=localhost user=nuradi  dbname=soloProject port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("ошибка при подключении к базе данных")
	}

	if err := db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Borrowing{}, &models.Reader{}); err != nil {

		panic("паника осторожно миграция не удалоось")
	}

	DB = db

	return db
}
