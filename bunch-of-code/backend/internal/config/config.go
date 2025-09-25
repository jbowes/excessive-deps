package config

import (
	"fmt"
	"os"
)

type Config struct {
	Addr string
	Env  string
}

func env(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func Load() Config {
	port := env("PORT", "8080")
	addr := fmt.Sprintf(":%s", port)
	return Config{
		Addr: addr,
		Env:  env("APP_ENV", "dev"),
	}
}
