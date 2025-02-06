package main

import (
	"fmt"
	"io"
	"log"

	"resty.dev/v3"
)

func main() {
	endpoint := "http://127.0.0.1:8080/"
	client := resty.New()
	resp, err := client.R().SetBody("http://my.net").Post(endpoint)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(string(body))
}
