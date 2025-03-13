package urlobject

import (
	"fmt"
	"math/rand/v2"
	"net/url"

	"github.com/google/uuid"
)

type URLObject struct {
	ID        string `json:"id"`
	OriginURL string `json:"originURL"`
	ShortURL  string `json:"shortURL"`
}

const allowedSimbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// NewURLObject create a URLObject with BaseURL=s and generate ShortURL
func NewURLObject(s string) (*URLObject, error) {
	u, err := url.Parse(s)
	if err != nil || u.Host == "" || u.Scheme == "" {
		return nil, fmt.Errorf("error validate url")
	}
	res := URLObject{}
	res.ID = uuid.NewString()
	res.OriginURL = s
	res.generateShortURL(7)

	return &res, nil
}

func (u *URLObject) generateShortURL(length int) {
	if u.OriginURL == "" {
		return
	}

	pass := make([]rune, length)
	for i := range pass {
		randomInt := rand.IntN(len(allowedSimbols) - 1)
		pass[i] = []rune(allowedSimbols)[randomInt]
	}
	u.ShortURL = string(pass)
}
