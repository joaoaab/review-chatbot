package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"review-bot/internal/config"
	"review-bot/internal/domain"
	"review-bot/internal/infrastructure/http"
	"review-bot/internal/infrastructure/http/controllers"
	"review-bot/internal/infrastructure/http/routers"
	"review-bot/internal/infrastructure/repository"
	"review-bot/internal/infrastructure/repository/bot_log"
	"review-bot/internal/infrastructure/repository/customer"
	"review-bot/internal/state"
)

func main() {
	uuid.EnableRandPool()

	err := godotenv.Load()
	if err != nil {
		log.Warn("Coudn't load .env file")
	}

	app := fx.New(
		config.Module,
		domain.Module,
		customer.Module,
		bot_log.Module,
		state.Module,
		repository.Module,
		controllers.Module,
		routers.Module,
		http.Module,
		fx.Invoke(func(server *fasthttp.Server) {}),
	)
	app.Run()
}
