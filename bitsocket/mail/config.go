package mail

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Server   string
	Port     int
	Email    string `env:"WATCHCASHEMAIL"`
	Password string `env:"PASSWORD"`
}

func (c *Config) Load() {
	err := env.Parse(c)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	c.Server = "smtp.gmail.com"
	c.Port = 587
}
