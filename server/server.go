package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"

	// "github.com/SkaisgirisMarius/order-processor.git/health"
	"github.com/SkaisgirisMarius/order-processor.git/order"
)

// NewRouter returns a new HTTP handler that implements the main server routes
func NewRouter(db *gorm.DB) http.Handler {

	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST"},
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler)
	r.Mount("/api/order", order.InitOrderRouter(db))
	// r.Mount("/api/health", health.InitRouter())
	return r
}
