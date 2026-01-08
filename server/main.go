package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ASHUTOSH-SWAIN-GIT/maturin/config"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/database"
	"github.com/ASHUTOSH-SWAIN-GIT/maturin/handlers"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	//  Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	//  Open Database Connection
	log.Println("Connecting to Supabase...")
	conn, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//  Check Connection
	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")

	// Initialize sqlc Queries
	queries := database.New(conn)

	//  Initialize Router
	router := handlers.NewRouter(queries)

	//  Start Server
	log.Printf("Server listening on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
