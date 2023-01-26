package loader

import (
	"github.com/durmusrasit/kampus-gql/models"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) (*DB, error) {
	err := AutoMigrate(db)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&models.Post{}); err != nil {
		return err
	}

	return nil
}
