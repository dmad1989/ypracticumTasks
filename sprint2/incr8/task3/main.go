package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", defaultHandle)
	mux.HandleFunc("/", LengthHandle)
	http.ListenAndServe(":3000", mux)
}

func defaultHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<html><body>"+strings.Repeat("Hello, world<br>", 20)+"</body></html>")
}

// LengthHandle возвращает размер распакованных данных.
func LengthHandle(w http.ResponseWriter, r *http.Request) {
	// создаём *gzip.Reader, который будет читать тело запроса
	// и распаковывать его
	var reader io.Reader
	if r.Header.Get("Accept-Encoding") == "gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	// при чтении вернётся распакованный слайс байт
	body, err := io.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Length: %d", len(body))
}
