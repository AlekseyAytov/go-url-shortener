package compress

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type compressWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w compressWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// если клиент отправил сжатые данные, то подменить request body
		if e := r.Header.Get("Content-Encoding"); strings.Contains(e, "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer gz.Close()
			r.Body = gz
		}

		// если клиент не поддерживает сжатие, то обработать стандартно
		if e := r.Header.Get("Accept-Encoding"); e == "" {
			next.ServeHTTP(w, r)
			return
		}

		if c := r.Header.Get("Content-Type"); c != "application/json" && c != "text/html" {
			next.ServeHTTP(w, r)
			return
		}

		// создаём gzip.Writer поверх текущего w
		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		// передаём обработчику страницы переменную типа gzipWriter для вывода данных
		next.ServeHTTP(compressWriter{ResponseWriter: w, Writer: gz}, r)
	})
}
