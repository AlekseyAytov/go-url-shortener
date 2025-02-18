package app

import (
	"encoding/json"
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
	middls []func(http.Handler) http.Handler
}

type _url struct {
	URL string `json:"url"`
}

type _result struct {
	Result string `json:"result"`
}

// NewShortenerAPI constructs a new ShortenerAPI,
// ensuring that the dependencies are valid values
func NewShortenerAPI(v *Vault, b string, middleWares []func(http.Handler) http.Handler) *ShortenerAPI {
	if v == nil || b == "" {
		panic("nil Vault!")
	}
	api := ShortenerAPI{
		vault:  v,
		router: mux.NewRouter(),
		base:   b,
		middls: middleWares,
	}
	api.endpoints()
	return &api
}

func (sh *ShortenerAPI) Router() *mux.Router {
	return sh.router
}

func (sh *ShortenerAPI) endpoints() {
	for _, m := range sh.middls {
		sh.router.Use(m)
	}
	sh.router.HandleFunc("/api/shorten", sh.shortURLJSON).Methods(http.MethodPost)
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
	ans := sh.base + "/" + obj.ShortURL
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(ans))
}

func (sh *ShortenerAPI) shortURLJSON(res http.ResponseWriter, req *http.Request) {
	var u _url
	if err := json.NewDecoder(req.Body).Decode(&u); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	objPtr, err := NewURLObject(u.URL)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	sh.vault.Add(*objPtr)

	r := _result{
		Result: sh.base + "/" + objPtr.ShortURL,
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(res).Encode(&r); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
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
