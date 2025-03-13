package main

import (
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
	"github.com/AlekseyAytov/go-url-shortener/internal/compress"
	"github.com/AlekseyAytov/go-url-shortener/internal/config"
	"github.com/AlekseyAytov/go-url-shortener/internal/logger"
	"github.com/AlekseyAytov/go-url-shortener/internal/storage/filestorage"
	"github.com/AlekseyAytov/go-url-shortener/internal/urlobject"
	"go.uber.org/zap"
)

func main() {
	l := logger.Get("Info")
	c := config.LoadOptions()
	// fmt.Printf("%q\n%q\n", c.BaseURL, c.SrvAdress)

	mw := []func(http.Handler) http.Handler{
		logger.RequestLogger,
		compress.GzipMiddleware,
	}

	storage := filestorage.NewFileStorage(c.StoragePath)
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
