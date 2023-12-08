package bot_log

import "go.uber.org/fx"

var Module = fx.Provide(NewBotLogRepository)
