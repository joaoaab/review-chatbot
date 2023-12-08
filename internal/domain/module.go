package domain

import "go.uber.org/fx"

var Module = fx.Provide(NewBotService)
