package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v11"
)

// Options является структурой для парсинга настроек из переменных окружения
type Options struct {
	SrvAdress   string `env:"SERVER_ADDRESS"`
	BaseURL     string `env:"BASE_URL"`
	StoragePath string `env:"FILE_STORAGE_PATH"`
}

// LoadOptions пробует:
// - спарсить настройки из переменных окружения
// - пустые значения взять из флагов
// - если значений нет, назначить по умолчанию
func LoadOptions() *Options {
	cfg := Options{}

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()

	if cfg.SrvAdress == "" {
		cfg.SrvAdress = flagServerSocket
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = flagServerBaseURL
	}

	if cfg.StoragePath == "" {
		cfg.StoragePath = flagFileStoragePath
	}
	return &cfg
}
