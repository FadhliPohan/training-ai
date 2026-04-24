package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"insightflow/be-penjualan/config"
)

// Pool is the global PostgreSQL connection pool.
var Pool *pgxpool.Pool

// Connect initialises the PostgreSQL connection pool using the DATABASE_URL from config.
func Connect() {
	cfg, err := pgxpool.ParseConfig(config.App.DatabaseURL)
	if err != nil {
		log.Fatalf("[database] failed to parse DATABASE_URL: %v", err)
	}

	// Connection pool tuning
	cfg.MaxConns = 20
	cfg.MinConns = 2
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("[database] failed to create connection pool: %v", err)
	}

	// Verify connectivity
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("[database] failed to ping PostgreSQL: %v", err)
	}

	Pool = pool
	log.Println("[database] connected to PostgreSQL successfully")
}

// Close gracefully closes the connection pool.
func Close() {
	if Pool != nil {
		Pool.Close()
		log.Println("[database] connection pool closed")
	}
}
