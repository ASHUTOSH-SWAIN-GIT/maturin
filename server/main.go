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
	// 1. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Open Database Connection
	log.Println("Connecting to Supabase...")
	conn, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 3. Check Connection
	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")

	// 4. Initialize sqlc Queries
	// 'queries' contains all your type-safe methods like CreateBucket, ListBuckets
	queries := database.New(conn)

	// 5. Initialize Router
	// Note: We need to update handlers.NewRouter to accept *database.Queries
	router := handlers.NewRouter(queries)

	// 6. Start Server
	log.Printf("Server listening on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
