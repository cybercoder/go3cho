package migrations

import (
	"3cho/config"
	"3cho/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func InitMigrate() {
	fmt.Println("1")
	DB, err := config.SetupDB()
	if err != nil {
		panic(err)
	}
	Migrate(DB)
}

func Migrate(db *gorm.DB) {
	fmt.Println("Migrating...")
	log.Println("Migrating")
	db.AutoMigrate(&models.User{})
	log.Println("Migrated")
	fmt.Println("Migrated...")
}
