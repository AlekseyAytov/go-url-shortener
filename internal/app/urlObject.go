package app

import (
	"fmt"
	"math/rand/v2"
	"net/url"
)

type URLObject struct {
	BaseURL  string
	ShortURL string
}

const allowedSimbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// NewURLObject create a URLObject with BaseURL=s and generate ShortURL
func NewURLObject(s string) (*URLObject, error) {
	u, err := url.Parse(s)
	if err != nil || u.Host == "" || u.Scheme == "" {
		return nil, fmt.Errorf("error validate url")
	}
	res := URLObject{}
	res.BaseURL = s
	res.generateShortURL(7)

	return &res, nil
}

func (u *URLObject) generateShortURL(length int) {
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
