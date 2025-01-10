package app

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HandlerContext struct {
	vault *Vault
}

// NewHandlerContext constructs a new HandlerContext,
// ensuring that the dependencies are valid values
func NewHandlerContext(v *Vault) *HandlerContext {
	if v == nil {
		panic("nil MongoDB session!")
	}
	return &HandlerContext{
		vault: v,
	}
}

func (ctx *HandlerContext) MainHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		body, _ := io.ReadAll(req.Body)
		long := string(body)
		obj, err := NewURLObject(long)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx.vault.Add(*obj)

		ans := "http://" + req.Host + "/" + obj.ShortURL
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(ans))
	case http.MethodGet:
		id := req.URL.Path
		id = strings.TrimLeft(id, "/")
		u, ok := ctx.vault.Find(id, func(u URLObject, s string) bool {
			return strings.Contains(u.ShortURL, s)
		})
		if ok {
			fmt.Println(u.BaseURL)
			res.Header().Set("Location", u.BaseURL)
			res.WriteHeader(http.StatusTemporaryRedirect)
		}
	default:
		res.WriteHeader(http.StatusBadRequest)
	}
}
