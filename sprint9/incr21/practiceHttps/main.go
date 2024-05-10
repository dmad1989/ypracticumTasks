package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello HTTPS!"))
	})
	// конструируем менеджер TLS-сертификатов
	manager := &autocert.Manager{
		// директория для хранения сертификатов
		Cache: autocert.DirCache("cache-dir"),
		// функция, принимающая Terms of Service издателя сертификатов
		Prompt: autocert.AcceptTOS,
		// перечень доменов, для которых будут поддерживаться сертификаты
		HostPolicy: autocert.HostWhitelist("mysite.ru", "www.mysite.ru"),
	}
	// конструируем сервер с поддержкой TLS
	server := &http.Server{
		Addr:    ":443",
		Handler: mux,
		// для TLS-конфигурации используем менеджер сертификатов
		TLSConfig: manager.TLSConfig(),
	}
	server.ListenAndServeTLS("", "")
}
