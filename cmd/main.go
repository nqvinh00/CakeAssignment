package main

import (
	"flag"

	"github.com/nqvinh00/CakeAssignment/model"
	"github.com/nqvinh00/CakeAssignment/pkg"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	logLevel   = flag.Int("log_level", 0, "Log level")
	configPath = flag.String("config", "config.yml", "Configuration file path")
)

func main() {
	flag.Parse()

	pkg.SetupLog(zerolog.Level(*logLevel))

	config := &model.Config{}
	if err := pkg.LoadFromYamlFile(*configPath, config); err != nil {
		log.Fatal().Str("path", *configPath).Err(err).Msg("failed to load config from file")
	}

	gormDB, err := pkg.ConnectDB(config.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

}
