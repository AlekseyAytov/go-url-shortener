package app

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Logging(log io.Writer, f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(log, "[%s]", time.Now().Format(time.DateTime))
		f(w, r)
	}
}
