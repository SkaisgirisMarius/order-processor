package server

import (
	"github.com/SkaisgirisMarius/order-processor.git/health"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"

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
	r.Mount("/api/health", health.InitHealthRouter())
	return r
}

func StartServer(handler http.Handler) {
	log.Println("Starting server on port :3000")
	httpSrv := makeHTTPServer(handler)
	httpSrv.Addr = ":3000"
	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Errorf("httpSrv.ListenAndServe() failed with %s", err)
	}
}

func makeHTTPServer(handler http.Handler) *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})

	httpsServer := makeServerFromMux(mux)
	httpsServer.Addr = "3000"

	return httpsServer
}

func makeServerFromMux(mux *http.ServeMux) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}
