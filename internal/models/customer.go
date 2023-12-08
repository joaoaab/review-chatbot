package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"review-bot/internal/structs"
)

type CustomerInformation struct {
	gorm.Model
	Name   string
	Email  string
	Phone  string
	Review CustomerReview
}

type CustomerReview struct {
	gorm.Model
	ProductName           string
	CustomerInformationID uint
	Log                   ChatbotLog
	Rating                int
}

type ChatbotLog struct {
	gorm.Model
	CustomerReviewID uint
	State            string
	Messages         datatypes.JSONSlice[structs.Message]
}
