package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi"))
	})
	r.Get("/item/{id}", func(rw http.ResponseWriter, r *http.Request) {
		// получаем значение URL-параметра id
		id := chi.URLParam(r, "id")
		io.WriteString(rw, fmt.Sprintf("item = %s", id))
	})
	// r передаётся как http.Handler
	http.ListenAndServe(":8080", r)
}
