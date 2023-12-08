package repository

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"review-bot/internal/config"
	"review-bot/internal/models"
)

func NewDb(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.CustomerInformation{}, &models.CustomerReview{}, &models.ChatbotLog{}); err != nil {
		panic(err)
	}

	log.Info("Migrated database")
	return db
}
