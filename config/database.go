package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	dsn := "user=postgres password=postgres dbname=mis_dev port=5432 sslmode=disable TimeZone=Asia/Tehran"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return DB, nil
}
