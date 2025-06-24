package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/arnokay/arnobot-shared/pkg/assert"
)

const (
	ENV_PORT   = "PORT"
	ENV_MB_URL = "MB_URL"
)

type config struct {
	Global GlobalConfig
	MB     MBConfig
}

type GlobalConfig struct {
	Env      string
	Port     int
	LogLevel int
}

type MBConfig struct {
	URL string
}

var Config *config

func Load() *config {
	Config = &config{
		Global: GlobalConfig{
			Port:     3000,
			LogLevel: -4,
		},
	}

	if os.Getenv(ENV_PORT) != "" {
		port, err := strconv.Atoi(os.Getenv(ENV_PORT))
		assert.NoError(err, fmt.Sprintf("%v: not a number", ENV_PORT))
		Config.Global.Port = port
	}

	flag.StringVar(&Config.Global.Env, "env", "development", "Environment (development|staging|production)")

	flag.IntVar(&Config.Global.Port, "port", Config.Global.Port, "Server Port")
	flag.IntVar(&Config.Global.LogLevel, "log-level", Config.Global.LogLevel, "Minimal Log Level (default: -4)")

	flag.StringVar(&Config.MB.URL, "mb-url", os.Getenv(ENV_MB_URL), "Message Broker URL")

	flag.Parse()

	return Config
}
