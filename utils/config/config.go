package config

import (
	"fmt"
	"os"
)

const (
	VERSION     string = "VERSION"
	DB_URL      string = "DB_URL"
	AUTH_SECRET string = "AUTH_SECRET"
)

var envVariables = []string{
	VERSION,
	DB_URL,
	AUTH_SECRET,
}

type Config struct {
	values map[string]string
}

func (c *Config) GetValue(key string) string {
	return c.values[key]
}

func GetConfig() *Config {
	cfg := &Config{
		values: map[string]string{},
	}

	for _, it := range envVariables {
		value, ok := os.LookupEnv(it)

		if !ok {
			panic(fmt.Sprintf("Missing %s", it))
		}

		cfg.values[it] = value
	}

	return cfg
}
