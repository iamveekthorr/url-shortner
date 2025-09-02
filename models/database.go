// Package models provides - Database functionality
package models

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ConnPool *pgxpool.Pool

func CloseDatabase() {
	if ConnPool != nil {
		ConnPool.Close()
		log.Println("Database connection closed")
	}
}

func InitDatabase(connectionString string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// TODO: IMPLEMENT CONNECTION RETRY LOGIC ON STARTUP
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	// ping to ensure DB is reachable at startup
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Println("✅ Connected to Database!")

	ConnPool = pool

	return nil
}
