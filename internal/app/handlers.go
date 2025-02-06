package app

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var ErrNotFoundValue = errors.New("value not found")

type ShortenerAPI struct {
	vault  *Vault
	router *mux.Router
	base   string
}

// NewShortenerAPI constructs a new ShortenerAPI,
// ensuring that the dependencies are valid values
func NewShortenerAPI(v *Vault, b string) *ShortenerAPI {
	if v == nil || b == "" {
		panic("nil Vault!")
	}
	api := ShortenerAPI{
		vault:  v,
		router: mux.NewRouter(),
		base:   b,
	}
	api.endpoints()
	return &api
}

func (sh *ShortenerAPI) Router() *mux.Router {
	return sh.router
}

func (sh *ShortenerAPI) endpoints() {
	// sh.router.Use(Logging)
	sh.router.HandleFunc("/{id}", sh.originalURL).Methods(http.MethodGet)
	sh.router.HandleFunc("/", sh.shortURL).Methods(http.MethodPost)
}

func (sh *ShortenerAPI) shortURL(res http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	long := string(body)
	obj, err := NewURLObject(long)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	sh.vault.Add(*obj)

	ans := "http://" + req.Host + "/" + obj.ShortURL
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(ans))
}

func (sh *ShortenerAPI) originalURL(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		http.Error(res, ErrNotFoundValue.Error(), http.StatusInternalServerError)
	}

	u, ok := sh.vault.Find(id, func(u URLObject, s string) bool {
		return strings.Contains(u.ShortURL, s)
	})
	if ok {
		res.Header().Set("Location", u.OriginURL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		http.Error(res, ErrNotFoundValue.Error(), http.StatusInternalServerError)
	}
}
