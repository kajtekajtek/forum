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

func QueryAllServers(db *gorm.DB) ([]models.Server, error) {
	var servers []models.Server

	if err := db.Find(&servers).Error; err != nil {
		return nil, fmt.Errorf("query all servers: %w", err)
	}

	return servers, nil
}

func QueryUserServers(db *gorm.DB, userID string) ([]models.Server, error) {
	var servers []models.Server

	// query user's memberships from database
	var memberships []models.Membership
	err := db.Where("user_id = ?", userID).Find(&memberships).Error
	if err != nil {
		return nil, fmt.Errorf("query user memberships: %w", err)
	}

	// get server IDs from memberships
	serverIDs := make([]uint, 0, len(memberships))
	for _, m := range memberships {
		serverIDs = append(serverIDs, m.ServerID)
	}

	// query servers with given IDs from database
	if len(serverIDs) > 0 {
		err := db.Where("id IN ?", serverIDs).Find(&servers).Error
		if err != nil {
		return nil, fmt.Errorf("query user servers: %w", err)
		}
	}

	return servers, nil
}
