package config

import (
	setup "chatgpt-service/pkg/errors"
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

func LoadConfig(filePath string, env string) (*GlobalConfig, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider(filePath), yaml.Parser()); err != nil {
		log.Printf("could not load config file: %v", err)
		return nil, fmt.Errorf("error loading config file: %v", err)
	}

	var cfg GlobalConfig
	if err := k.Unmarshal(fmt.Sprintf("%s", env), &cfg); err != nil {
		log.Printf("could not unmarshal config file: %v", err)
		return nil, fmt.Errorf("error unmarshaling config file: %v", err)
	}
	cfg.OpenAIEnv.ParseEnv(k, env)

	if cfg.Environment == "" {
		logrus.WithFields(logrus.Fields{
			"component": "setup",
			"env":       env,
		}).Error(logrus.ErrorLevel, "Failed to load config")
		return nil, setup.LoadConfigError()
	}

	return &cfg, nil
}
