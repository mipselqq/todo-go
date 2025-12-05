package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	host          string
	port          uint16
	willServeDocs bool
}

func FromEnv() (Config, error) {
	host := os.Getenv("HOST")
	if host == "" {
		return Config{}, fmt.Errorf("HOST is required")
	}

	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil {
		return Config{}, fmt.Errorf("PORT is required and must be a number")
	}

	return Config{
		host:          host,
		port:          uint16(port),
		willServeDocs: os.Getenv("ENV") != "production",
	}, nil
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func (c Config) Host() string {
	return c.host
}

func (c Config) Port() uint16 {
	return c.port
}

func (c Config) WillServeDocs() bool {
	return c.willServeDocs
}
