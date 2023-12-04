package main

import (
	"flag"
	"fmt"
	"github.com/stefanicai/transact/internal/app"
	"github.com/stefanicai/transact/internal/config"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configFilePath := flag.String("configFile", "./config/local.yaml", "Configuration of the application.")
	flag.Parse()
	cfg, err := config.LoadFromFile(*configFilePath)
	if err != nil {
		slog.Error(fmt.Sprintf("could not parse config file %s", *configFilePath), err)
		os.Exit(1)
	}

	//start app
	app.Start(cfg)
}
