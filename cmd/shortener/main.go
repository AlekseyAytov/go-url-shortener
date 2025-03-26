package main

import (
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
	"github.com/AlekseyAytov/go-url-shortener/internal/compress"
	"github.com/AlekseyAytov/go-url-shortener/internal/config"
	"github.com/AlekseyAytov/go-url-shortener/internal/logger"
	"github.com/AlekseyAytov/go-url-shortener/internal/storage/db"
	"github.com/AlekseyAytov/go-url-shortener/internal/storage/filestorage"
	"github.com/AlekseyAytov/go-url-shortener/internal/urlobject"
	"go.uber.org/zap"
)

func main() {
	l := logger.Get("Info")
	c, err := config.LoadOptions()
	if err != nil {
		l.Fatal("required only one storage method:",
			zap.String("fileStoragePath", c.StoragePath),
			zap.String("databaseDSN", c.DatabaseDSN),
		)
	}
	l.Info(
		"loaded configurtion: ",
		zap.String("socket", c.SrvAdress),
		zap.String("baseURL", c.BaseURL),
		zap.String("filePath", c.StoragePath),
		zap.String("dbDSN", c.DatabaseDSN),
	)

	mw := []func(http.Handler) http.Handler{
		logger.RequestLogger,
		compress.GzipMiddleware,
	}

	storage := getStorage(c, l)

	v := urlobject.GetVault(storage)
	api := app.NewShortenerAPI(v, c.BaseURL, mw)

	l.Info(
		"starting application server on socket: "+c.SrvAdress,
		zap.String("socket", c.SrvAdress),
	)

	l.Fatal(
		"server closed",
		zap.Error(http.ListenAndServe(c.SrvAdress, api.Router())),
	)
}

func getStorage(c *config.Options, l *zap.Logger) urlobject.PersistentStorage {
	if c.StoragePath != "" {
		return filestorage.NewFileStorage(c.StoragePath)
	}

	s, err := db.NewDBStorage(c.DatabaseDSN)
	if err != nil {
		l.Fatal("error connection to db",
			zap.Error(err))
	}
	return s
}
