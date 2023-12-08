package customer

import (
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"review-bot/internal/models"
	"review-bot/internal/structs"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (cr *CustomerRepository) CreateCustomer(customer models.CustomerInformation) (models.CustomerInformation, error) {
	err := cr.db.Create(&customer).Error
	return customer, err
}

func (cr *CustomerRepository) GetCustomer(id uint) (models.CustomerInformation, error) {
	var customer models.CustomerInformation
	err := cr.db.Preload(clause.Associations).Where("id = ?", id).First(&customer).Error
	return customer, err
}

func (cr *CustomerRepository) StartReview(id uint, productName string, startState string) (models.CustomerReview, error) {
	startReviewMessage := structs.Message{
		IsUserMessage: false,
		Text:          fmt.Sprintf("Hello again! We noticed you've recently received your %s. We'd love to hear about\nyour experience. Can you spare a few minutes to share your thoughts?", productName),
	}

	review := models.CustomerReview{
		ProductName:           productName,
		CustomerInformationID: id,
		Log: models.ChatbotLog{
			State:    startState,
			Messages: datatypes.NewJSONSlice[structs.Message]([]structs.Message{startReviewMessage}),
		},
		Rating: 0,
	}

	cr.db.Create(&review)
	return review, nil
}

func (cr *CustomerRepository) GetReview(id uint) (models.CustomerReview, error) {
	var response models.CustomerReview
	err := cr.db.Preload(clause.Associations).Where("id = ?", id).First(&response).Error
	return response, err
}

func (cr *CustomerRepository) SetReviewScore(customerReviewID uint, score int) {
	cr.db.Model(&models.CustomerReview{}).Where("id = ?", customerReviewID).Update("rating", score)
}
