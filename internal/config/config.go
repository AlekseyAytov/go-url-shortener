package config

import (
	"errors"
	"flag"
	"log"

	"github.com/caarlos0/env/v11"
)

var ErrNotSingleStorage = errors.New("required only one storage method")

// Options является структурой для парсинга настроек из переменных окружения
type Options struct {
	SrvAdress   string `env:"SERVER_ADDRESS"`
	BaseURL     string `env:"BASE_URL"`
	StoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN string `env:"DATABASE_DSN"`
}

// LoadOptions пробует:
// - спарсить настройки из переменных окружения
// - пустые значения взять из флагов
// - если значений нет, назначить по умолчанию
func LoadOptions() (*Options, error) {
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

	if cfg.DatabaseDSN == "" {
		cfg.DatabaseDSN = flagDatabaseDSN
	}

	err = cfg.checkConfig()
	if err != nil {
		return &cfg, err
	}
	cfg.fillDefaultValues()
	return &cfg, nil
}

func (o *Options) checkConfig() error {
	if o.StoragePath != "" && o.DatabaseDSN != "" {
		return ErrNotSingleStorage
	}
	return nil
}

func (o *Options) fillDefaultValues() {
	if o.SrvAdress == "" {
		o.SrvAdress = ":8080"
	}
	if o.BaseURL == "" {
		o.BaseURL = "http://localhost:8080"
	}
	if o.StoragePath == "" && o.DatabaseDSN == "" {
		o.StoragePath = "/tmp/short-url-db.json"
	}
	// if o.DatabaseDSN == "" {
	// 	o.DatabaseDSN = "postgres://supportuser:1234@localhost:5432/demo"
	// }
}
