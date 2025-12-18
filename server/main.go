package main

import (
	"log"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/config"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
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

	log.Println("Successfully connected to the database!")
}
