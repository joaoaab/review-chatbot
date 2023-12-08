package bot_log

import (
	"gorm.io/gorm"
	"review-bot/internal/models"
	"review-bot/internal/structs"
)

type BotLogRepository struct {
	db *gorm.DB
}

func NewBotLogRepository(db *gorm.DB) *BotLogRepository {
	return &BotLogRepository{db: db}
}

func (blr *BotLogRepository) ChangeState(customerReviewID uint, state string) {
	blr.db.Model(&models.ChatbotLog{}).Where("customer_review_id = ?", customerReviewID).Update("state", state)
}

func (blr *BotLogRepository) AppendMessage(customerReviewID uint, msg structs.Message) {
	var log models.ChatbotLog
	blr.db.Where("customer_review_id = ?", customerReviewID).First(&log)
	messages := log.Messages
	messages = append(messages, msg)
	blr.db.Model(&models.ChatbotLog{}).Where("customer_review_id = ?", customerReviewID).Update("messages", messages)
}

func (blr *BotLogRepository) RetrieveMessages(customerReviewID uint) []structs.Message {
	var chatlog models.ChatbotLog
	blr.db.Where("customer_review_id = ?", customerReviewID).First(&chatlog)
	return chatlog.Messages
}
