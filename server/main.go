package main

import (
	"log"
	"net/http"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/config"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to Supabase...")
	store, err := database.NewStore(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()
	log.Println("Database connection established")

	// Initialize Router
	router := handlers.NewRouter(store)

	log.Printf("Server listening on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
