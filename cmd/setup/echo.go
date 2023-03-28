package setup

import (
	"chatgpt-service/internal/api"
	"chatgpt-service/internal/config"
	"chatgpt-service/internal/pkg/client"
	"chatgpt-service/internal/pkg/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitializeEcho(cfg *config.GlobalConfig, oc client.OpenAIClient, db store.Database) (error, *echo.Echo) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	//ConfigHandler(e, *cfg, oc)
	err := api.SetupRoutes(e, *cfg, oc, db)
	if err != nil {
		return err, nil
	}

	return nil, e
}
