package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var api *ShortenerAPI

func TestMain(m *testing.M) {
	v := GetVault()
	api = NewShortenerAPI(v)
	m.Run()
}

func TestShortenerAPI_shortURL(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name          string
		requestPath   string
		requestMethod string
		requestBody   string
		want          want
	}{
		{
			name:          "first, positive",
			requestPath:   "/",
			requestMethod: http.MethodPost,
			requestBody:   "http://ya.ru",
			want: want{
				code:        http.StatusCreated,
				response:    "http://",
				contentType: "text/plain",
			},
		},
		{
			name:          "second, negative",
			requestPath:   "/",
			requestMethod: http.MethodGet,
			requestBody:   "http://ya.ru",
			want: want{
				code: http.StatusMethodNotAllowed,
			},
		},
		{
			name:          "third, negative",
			requestPath:   "/",
			requestMethod: http.MethodPost,
			requestBody:   "htt:/ya.ru",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody := strings.NewReader(tt.requestBody)
			request := httptest.NewRequest(tt.requestMethod, tt.requestPath, reqBody)
			w := httptest.NewRecorder()
			api.Router().ServeHTTP(w, request)

			res := w.Result()

			// проверка статус кода
			assert.Equal(t, tt.want.code, res.StatusCode)

			// проверка заголовка
			assert.Contains(t, res.Header.Get("Content-Type"), tt.want.contentType)

			// проверка тела ответа
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Contains(t, string(resBody), tt.want.response)
		})
	}
}

func TestShortenerAPI_originalURL(t *testing.T) {
	obj1, err := NewURLObject("http://google.com")
	require.NoError(t, err)
	obj2, err := NewURLObject("http://ya.ru")
	require.NoError(t, err)

	api.vault.Add(*obj1)
	api.vault.Add(*obj2)

	type want struct {
		code     int
		location string
	}

	tests := []struct {
		name          string
		requestPath   string
		requestMethod string
		want          want
	}{
		{
			name:          "first, positive",
			requestPath:   "/" + obj1.ShortURL,
			requestMethod: http.MethodGet,
			want: want{
				code:     http.StatusTemporaryRedirect,
				location: obj1.BaseURL,
			},
		},
		{
			name:          "second, positive",
			requestPath:   "/" + obj2.ShortURL,
			requestMethod: http.MethodGet,
			want: want{
				code:     http.StatusTemporaryRedirect,
				location: obj2.BaseURL,
			},
		},
		{
			name:          "third, negative",
			requestPath:   "/" + "11Ghhh6",
			requestMethod: http.MethodGet,
			want: want{
				code:     http.StatusInternalServerError,
				location: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.requestMethod, tt.requestPath, nil)
			w := httptest.NewRecorder()
			api.Router().ServeHTTP(w, request)

			res := w.Result()

			// проверка статус кода
			assert.Equal(t, tt.want.code, res.StatusCode)
			// проверка заголовка
			assert.Contains(t, res.Header.Get("Location"), tt.want.location)
		})
	}
}
