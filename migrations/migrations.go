package migrations

import (
	"github.com/KirillEL/golang-postgres-rest-api/models"
	"gorm.io/gorm"
)

func MigrateCars(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Car{})
	return err
}
