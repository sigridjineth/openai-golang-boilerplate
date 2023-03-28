package main

import (
	"chatgpt-service/cmd/setup"
	"chatgpt-service/internal/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	errCh := make(chan error)

	go func() {
		port := 8080
		// load configuration
		env := os.Getenv("ENV")
		cfg, err := config.LoadConfig(config.DefaultConfigPath, env)
		if err != nil {
			errCh <- err
		}
		// setup database
		db := setup.InitializeDatabase(cfg)
		// setup openai client
		oc, err := setup.NewOpenAIClient(cfg)
		if err != nil {
			errCh <- err
		}
		// setup echo server
		err, e := setup.InitializeEcho(cfg, *oc, *db)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":   "chatgpt-service",
				"component": "main",
			}).WithError(err).Error("ChatGPT Service API server running failed")
			errCh <- err
		}
		fmt.Println("ChatGPT Service Server is now running at the port :" + fmt.Sprint(port))
		logrus.WithFields(logrus.Fields{
			"service":   "chatgpt-service",
			"component": "main",
			"port":      port,
		}).Info("ChatGPT Service API server running")
		err = e.Start(":" + fmt.Sprint(port))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":   "chatgpt-service",
				"component": "main",
			}).WithError(err).Error("ChatGPT Service API server running failed")
			errCh <- err
		}
	}()

	for {
		select {
		case err := <-errCh:
			logrus.WithFields(logrus.Fields{
				"service":   "chatgpt-service",
				"component": "main",
			}).WithError(err).Error("An error occurred")
			wg.Done()
			os.Exit(1)
		}
	}

	wg.Wait()
}
