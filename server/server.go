package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()
	r.Get("/", onIndex)

	return r
}

func onIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
