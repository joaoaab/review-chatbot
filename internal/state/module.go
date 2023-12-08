package state

import "go.uber.org/fx"

var Module = fx.Provide(NewStateMachineHandler)
