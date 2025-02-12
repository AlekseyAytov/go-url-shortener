package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v11"
)

type Options struct {
	SrvAdress string `env:"SERVER_ADDRESS"`
	BaseURL   string `env:"BASE_URL"`
}

func LoadOptions() *Options {
	cfg := Options{}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	srvAdressFlag := flag.String("a", ":8080", "server socket")
	baseURLflag := flag.String("b", "http://localhost:8080", "base address")
	flag.Parse()

	if cfg.SrvAdress == "" {
		cfg.SrvAdress = *srvAdressFlag
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = *baseURLflag
	}
	return &cfg
}
