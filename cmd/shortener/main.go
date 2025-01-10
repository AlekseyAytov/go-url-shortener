package main

import (
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
)

func main() {
	v := app.GetVault()
	hctx := app.NewHandlerContext(v)

	mux := http.NewServeMux()
	// mux.HandleFunc(`/`, app.Logging(os.Stdout, mainPage))
	mux.HandleFunc(`/`, hctx.MainHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
