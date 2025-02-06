package main

import (
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
	"github.com/AlekseyAytov/go-url-shortener/internal/config"
)

func main() {
	c := config.LoadOptions()
	// fmt.Printf("%q\n%q\n", c.BaseURL, c.SrvAdress)

	v := app.GetVault()
	api := app.NewShortenerAPI(v, c.BaseURL)

	err := http.ListenAndServe(c.SrvAdress, api.Router())
	if err != nil {
		panic(err)
	}
}
