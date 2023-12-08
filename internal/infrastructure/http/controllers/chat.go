package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"review-bot/internal/domain"
	"review-bot/internal/infrastructure/repository/customer"
	"review-bot/internal/models"
	"review-bot/internal/state"
	"strconv"
	"time"
)

type ChatController struct {
	customerRepository *customer.CustomerRepository
	BotService         *domain.BotService
}

func NewChatController(repository *customer.CustomerRepository, service *domain.BotService) *ChatController {
	return &ChatController{customerRepository: repository, BotService: service}
}

func (cc *ChatController) Signup(c *fiber.Ctx) error {
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	name := c.FormValue("name")

	customer := models.CustomerInformation{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	customer, err := cc.customerRepository.CreateCustomer(customer)
	if err != nil {
		log.Fatal(err)
	}

	idCookie := new(fiber.Cookie)
	idCookie.Name = "id"
	idCookie.Path = "/"
	idCookie.Expires = time.Now().Add(24 * time.Hour)
	idCookie.Value = strconv.Itoa(int(customer.ID))
	c.Cookie(idCookie)
	c.Set("HX-Location", "http://localhost:8080/chat")
	return c.SendStatus(http.StatusOK)
}

func (cc *ChatController) StartChat(c *fiber.Ctx) error {
	payload := struct {
		ID          uint   `json:"id"`
		ProductName string `json:"product_name"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	_, err := cc.customerRepository.StartReview(payload.ID, payload.ProductName, state.StateWaitingSentiment)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}

func (cc *ChatController) ChatMessage(c *fiber.Ctx) error {
	message := c.FormValue("message")
	id := GetIdFromCookies(c)
	err := cc.BotService.HandleMessage(message, id)
	if err != nil {
		return err
	}
	return c.SendStatus(http.StatusOK)
}
