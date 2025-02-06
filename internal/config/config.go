package config

import "flag"

type Options struct {
	SrvAdress string
	BaseURL   string
}

func LoadOptions() *Options {
	res := Options{}
	flag.StringVar(&res.SrvAdress, "a", ":8080", "server socket")
	flag.StringVar(&res.BaseURL, "b", "localhost:8080", "base address")
	flag.Parse()
	return &res
}
