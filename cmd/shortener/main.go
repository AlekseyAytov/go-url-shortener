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
		res.Header().Set("content-type", "text/plain")
		body, _ := io.ReadAll(req.Body)
		fmt.Println(string(body))
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}
