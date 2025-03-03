package compress

import (
	"compress/gzip"
	"io"
	"net/http"
)

type compressWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w compressWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func GzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if e := r.Header.Get("Accept-Encoding"); e == "" {
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
