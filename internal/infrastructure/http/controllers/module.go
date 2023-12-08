package controllers

import "go.uber.org/fx"

var Module = fx.Provide(NewChatController, NewUIController)
