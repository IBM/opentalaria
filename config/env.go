package config

import (
	"log"
	"log/slog"
	"strings"
)

type OTProfile int

const (
	Localdev OTProfile = iota
	Dev
	Prod
	Unknown
)

func (c *Config) loadProfile() {
	switch c.Env.GetString("profile") {
	case "localdev":
		c.OTProfile = Localdev
	case "dev":
		c.OTProfile = Dev
	case "prod":
		c.OTProfile = Prod
	default:
		c.OTProfile = Unknown
	}
}

func (c *Config) loadLogLevel() {
	switch strings.ToLower(c.Env.GetString("log.level")) {
	case "debug":
		c.LogLevel = slog.LevelDebug
	case "info":
		c.LogLevel = slog.LevelInfo
	case "warn":
		c.LogLevel = slog.LevelWarn
	case "error":
		c.LogLevel = slog.LevelError
	default:
		log.Println("no log level set or value is invalid, setting default WARN level")
		c.LogLevel = slog.LevelWarn
	}
}
