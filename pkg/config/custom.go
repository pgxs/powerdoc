package config

import (
	"os"
	cfg "pgxs.io/chassis/config"
)

const (
	envKey = "POWERDOC_CONF"
)

type ServerConfig struct {
	JWT JWTConfig `yaml:"jwt"`
}

type JWTConfig struct {
	JWTHmacSignSecret string `yaml:"hmac-sign-secret"`
}

var customConfig *ServerConfig

func Load() {
	cfg.SetLoadFileEnvKey(envKey)
	cfg.LoadFromEnvFile()
	customConfig = new(ServerConfig)
	cfg.LoadCustomFromFile(os.Getenv(envKey), customConfig)
}

func Server() *ServerConfig {
	if customConfig == nil {
		Load()
	}
	return customConfig
}
