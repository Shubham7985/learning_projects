package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Postgres driver
)

// Supabase is the global DB connection
var Supabase *sql.DB

func ConnectDB() {
	// Direct connection string (replace with your real password)
	dsn := "postgresql://postgres:sk%40%23%407985zU96gf@db.hqucaraqsltyjprprjjz.supabase.co:5432/postgres"

	var err error
	Supabase, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to Supabase:", err)
	}

	err = Supabase.Ping()
	if err != nil {
		log.Fatal("Database ping failed:", err)
	}

	fmt.Println("Connected to Supabase Postgres successfully!")
}
