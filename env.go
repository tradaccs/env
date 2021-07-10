package env

import (
	"os"
)

type Config struct {
	Vars []string
}

func WithVars(vars ...string) *Config {
	return &Config{Vars: vars}
}

type Env map[string]string

var env Env

func New(c *Config) Env {
	if env == nil {
		env = make(Env, len(c.Vars))
		env.fill(c.Vars)
	}
	return env
}

func (e Env) fill(keys []string) {
	for _, key := range keys {
		e[key] = os.Getenv(key)
	}
}

func FromFile(path string) {
}
