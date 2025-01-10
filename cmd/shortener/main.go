package main

import (
	"fmt"
	"io"
	"net/http"
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
		fmt.Println(string(body))

		// res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusOK)
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}
