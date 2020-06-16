package config

import (
	cfg "pgxs.io/chassis/config"
)

const envKey = "POWERDOC_CONF"

func Load() {
	cfg.SetLoadFileEnvKey(envKey)
	cfg.LoadFromEnvFile()
}
