package api

import (
	"chatgpt-service/internal/config"
	gclient "chatgpt-service/internal/pkg/client"
	"chatgpt-service/internal/pkg/store"
	"chatgpt-service/pkg/client"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, cfg config.GlobalConfig, oc gclient.OpenAIClient, db store.Database) error {
	hd, err := NewHandler(e.AcquireContext(), cfg, &oc, &db)
	if err != nil {
		return err
	}
	//e.Use(HttpRequestLogHandler) -- FIX THE ERROR
	e.GET(HealthEndpoint, HealthCheck)
	e.GET(client.GetAllModels, hd.ListModels)
	e.GET(client.RetrieveModels, hd.RetrieveModel)
	e.POST(client.CreateCompletionEndpoint, hd.CreateCompletion)

	return nil
}
