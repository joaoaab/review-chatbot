package routers

import (
	"github.com/gofiber/fiber/v2"
	"review-bot/internal/infrastructure/http/controllers"
)

type ChatRouter struct {
	controller *controllers.ChatController
}

func (cr *ChatRouter) Load(r *fiber.App) {
	r.Post("/api/signup", cr.controller.Signup)
	r.Post("/api/chat-message", cr.controller.ChatMessage)
	r.Post("/api/create-product-review", cr.controller.StartChat)
}

func NewChatRouter(controller *controllers.ChatController) *ChatRouter {
	return &ChatRouter{controller: controller}
}
