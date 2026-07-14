package routes

import (
	"net/http"

	"github.com/AdarshJha-1/Depot/internal/handler"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", handler.HandlerPong())
	mux.Handle("GET /", http.FileServer(http.Dir("./web")))

	mux.HandleFunc("POST /uploads", handler.HandlerUpload())
	mux.HandleFunc("GET /uploads/{img}", handler.HandlerGetUpload())

	return mux
}
