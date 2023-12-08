package controllers

import (
	"github.com/gofiber/fiber/v2"
	"review-bot/internal/infrastructure/repository/bot_log"
	"review-bot/internal/infrastructure/repository/customer"
)

type UIController struct {
	customerRepository *customer.CustomerRepository
	botLogRepository   *bot_log.BotLogRepository
}

func NewUIController(repo *customer.CustomerRepository, botLogRepo *bot_log.BotLogRepository) *UIController {
	return &UIController{customerRepository: repo, botLogRepository: botLogRepo}
}

func (uc *UIController) Signup(c *fiber.Ctx) error {
	return c.Render("signup", fiber.Map{})
}

func (uc *UIController) Chat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{})
}

func (uc *UIController) ChatLog(c *fiber.Ctx) error {
	customerID := GetIdFromCookies(c)
	customerModel, err := uc.customerRepository.GetCustomer(customerID)
	if err != nil {
		return err
	}
	logs := uc.botLogRepository.RetrieveMessages(customerModel.Review.ID)
	return c.Render("chat-messages", fiber.Map{"Messages": logs})
}
