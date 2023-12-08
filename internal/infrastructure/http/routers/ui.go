package routers

import (
	"github.com/gofiber/fiber/v2"
	"review-bot/internal/infrastructure/http/controllers"
)

type UIRouter struct {
	controller *controllers.UIController
}

func (ur *UIRouter) Load(r *fiber.App) {
	r.Get("/signup", ur.controller.Signup)
	r.Get("/chat", ur.controller.Chat)
	r.Get("/chat-log", ur.controller.ChatLog)
}

func NewUIRouter(controller *controllers.UIController) *UIRouter {
	return &UIRouter{controller: controller}
}
