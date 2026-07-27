package shortener

import (
	"time"
)

type Config struct {
	codeLength  int
	defaultTTL  time.Duration
	maxURLLen   int
	allowCustom bool
	validator   func(string) error
}

type Option func(*Config)

func defaultConfig() Config {
	return Config{
		codeLength:  8,
		defaultTTL:  24 * time.Hour,
		maxURLLen:   2048,
		allowCustom: false,
		validator:   ValidateURL,
	}
}

func WithCodeLength(n int) Option {
	return func(c *Config) {
		c.codeLength = n
	}
}

func WithTTL(d time.Duration) Option {
	return func(c *Config) {
		c.defaultTTL = d
	}
}

func WithCustomCodes() Option {
	return func(c *Config) {
		c.allowCustom = true
	}
}

func WithValidator(v func(string) error) Option {
	return func(c *Config) {
		c.validator = v
	}
}
