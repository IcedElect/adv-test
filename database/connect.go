package database

import (
	"icedelect/avito-test-adv/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=avito.adv"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.Adv{},
		&models.Photo{},
	)

	return db
}
