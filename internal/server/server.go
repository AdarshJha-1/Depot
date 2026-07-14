package server

import (
	"net/http"
	"os"
	"time"

	"github.com/AdarshJha-1/Depot/internal/middleware"
	"github.com/AdarshJha-1/Depot/internal/routes"
)

func New() *http.Server {

	r := routes.New()
	wrappedRouter := middleware.LogAPIInfo(middleware.SetResponseHeaders(middleware.RateLimit(r)))
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	svr := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           wrappedRouter,
	}

	return svr
}
