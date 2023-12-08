package customer

import "go.uber.org/fx"

var Module = fx.Provide(NewCustomerRepository)
