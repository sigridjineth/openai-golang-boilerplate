package setup

import (
	"chatgpt-service/internal/config"
	"chatgpt-service/internal/pkg/client"
	"github.com/labstack/echo/v4"
)

func ConfigHandler(e *echo.Echo, cfg config.GlobalConfig, oc *client.OpenAIClientInterface) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Set(config.GlobalConfigKey, cfg)
			ctx.Set(client.OpenAIClientKey, oc)
			return next(ctx)
		}
	})
}
