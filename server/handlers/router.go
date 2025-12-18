package handlers

import (
	"net/http"
	"os"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func NewRouter(store *database.Store) http.Handler {
	r := chi.NewRouter()
	_ = godotenv.Load()

	//middleware for  logging and crash recovery
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	//middleware to handle cors error
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("NEXT_URL")},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return r
}
