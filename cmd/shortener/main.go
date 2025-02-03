package main

import (
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
)

func main() {
	v := app.GetVault()
	api := app.NewShortenerAPI(v)

	err := http.ListenAndServe(`:8080`, api.Router())
	if err != nil {
		panic(err)
	}
}
