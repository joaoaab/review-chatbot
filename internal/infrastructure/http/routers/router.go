package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Router interface {
	Load()
}

func MakeRouter(chat *ChatRouter, ui *UIRouter) *fiber.App {
	engine := html.New("./cmd/templates", ".html")

	cfg := fiber.Config{
		AppName:       "Running Chatbot",
		CaseSensitive: true,
		Views:         engine,
	}

	r := fiber.New(cfg)

	chat.Load(r)
	ui.Load(r)

	return r
}
