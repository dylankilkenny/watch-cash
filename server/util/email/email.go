package email

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/caarlos0/env"
)

type config struct {
	Email    string `env:"WATCHCASHEMAIL"`
	Password string `env:"PASSWORD"`
}

func Send(subject string, body string, to string) {
	cfg := config{}
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	msg := "From: " + cfg.Email + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", cfg.Email, cfg.Password, "smtp.gmail.com"),
		cfg.Email, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
}
