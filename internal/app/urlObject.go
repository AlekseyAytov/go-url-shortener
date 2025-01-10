package app

import (
	"math/rand/v2"
	"net/url"
)

type URLObject struct {
	BaseURL  string
	ShortURL string
}

const allowedSimbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func NewURLObject(s string) (*URLObject, error) {
	res := URLObject{}

	_, err := url.Parse(s)
	if err != nil {
		return &res, err
	}

	res.BaseURL = s
	res.GenerateShortURL(7)

	return &res, nil
}

func (u *URLObject) GenerateShortURL(length int) {
	if u.BaseURL == "" {
		return
	}

	pass := make([]rune, length)
	for i := range pass {
		randomInt := rand.IntN(len(allowedSimbols) - 1)
		pass[i] = []rune(allowedSimbols)[randomInt]
	}
	u.ShortURL = string(pass)
}
