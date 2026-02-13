package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func Load() (Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if _, err := strconv.Atoi(port); err != nil {
		return Config{}, fmt.Errorf("invalid PORT: %w", err)
	}

	port = ":" + port

	readTimeout := os.Getenv("READ_TIMEOUT")

	writeTimeout := os.Getenv("WRITE_TIMEOUT")

	idleTimeout := os.Getenv("IDLE_TIMEOUT")

	return Config{
		Port:         port,
		ReadTimeout:  parseDuration(readTimeout, 5*time.Second),
		WriteTimeout: parseDuration(writeTimeout, 5*time.Second),
		IdleTimeout:  parseDuration(idleTimeout, 30*time.Second),
	}, nil
}

func parseDuration(s string, defaultValue time.Duration) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil || d <= 0 {
		return defaultValue
	}
	return d
}
