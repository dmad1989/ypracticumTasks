package main

import (
	"compress/gzip"
	"compress/zlib"
	"io"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandle)
	http.ListenAndServe(":3000", zlibHandle(mux))
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	// w.Writer будет отвечать за gzip-сжатие, поэтому пишем в него
	return w.Writer.Write(b)
}

func gzipHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем, что клиент поддерживает gzip-сжатие
		// это упрощённый пример. В реальном приложении следует проверять все
		// значения r.Header.Values("Accept-Encoding") и разбирать строку
		// на составные части, чтобы избежать неожиданных результатов
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// если gzip не поддерживается, передаём управление
			// дальше без изменений
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
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

type zlibWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w zlibWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func zlibHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "deflate") {
			next.ServeHTTP(w, r)
			return
		}
		zl, err := zlib.NewWriterLevel(w, zlib.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer zl.Close()
		w.Header().Set("Content-Encoding", "deflate")
		next.ServeHTTP(zlibWriter{ResponseWriter: w, Writer: zl}, r)
	})
}