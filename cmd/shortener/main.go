package main

import (
	"io"
	"net/http"

	"github.com/AlekseyAytov/go-url-shortener/internal/app"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainPage)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		body, _ := io.ReadAll(req.Body)
		long := string(body)
		obj, err := app.NewURLObject(long)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		ans := "http://" + req.Host + "/" + obj.ShortURL
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(ans))
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}
