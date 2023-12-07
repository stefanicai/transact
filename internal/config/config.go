package config

import (
	"github.com/go-faster/yaml"
	"github.com/stefanicai/transact/internal/forex"
	"github.com/stefanicai/transact/internal/persistence/mongodb"
	"log/slog"
	"os"
)

// Config holds configuration provided when the application is started
// At this point it's only the rates file location, but it'll likely grow. E.g. log level etc
type Config struct {
	Forex forex.Config   `yaml:"forex"`
	Mongo mongodb.Config `yaml:"mongo"`
}

func LoadFromFile(configFilePath string) (Config, error) {
	var cfg Config
	slog.Info("loading config", "file", configFilePath)
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}
	slog.Debug("config file read successfully", "file", configFilePath)

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return cfg, err
	}
	slog.Info("config successfully loaded from file", "file", configFilePath)
	return cfg, nil
}
