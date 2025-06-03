package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"github.com/kajtekajtek/forum/backend/internal/config"
	"github.com/kajtekajtek/forum/backend/internal/models"
)

func Initialize(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
	)
	
	// initialize database session
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("open database session: %w", err)
	}

	// run auto migration for application's models
	err = db.AutoMigrate(&models.Server{}, &models.Membership{}); 
	if err != nil {
		return &gorm.DB{}, fmt.Errorf("auto migrate: %w", err)
	}

	return db, nil
}
	
