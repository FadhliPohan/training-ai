package database

import (
	"context"
	"log"
	"time"

	"insightflow/be-penjualan/config"
	"insightflow/be-penjualan/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func RunAutoMigrate() {
	db, err := gorm.Open(postgres.Open(config.App.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("[database] auto-migrate failed: cannot open gorm connection: %v", err)
	}

	setupStatements := []string{
		`CREATE EXTENSION IF NOT EXISTS pgcrypto`,
		`CREATE SCHEMA IF NOT EXISTS app`,
		`CREATE SCHEMA IF NOT EXISTS bisnis`,
	}
	for _, stmt := range setupStatements {
		if execErr := db.Exec(stmt).Error; execErr != nil {
			log.Fatalf("[database] auto-migrate schema setup failed: %v", execErr)
		}
	}

	if err := db.AutoMigrate(
		&domain.User{},
		&domain.TelegramConfig{},
		&domain.AnomalyConfig{},
		&domain.SavedDashboard{},
		&domain.Produk{},
		&domain.Customer{},
		&domain.Order{},
		&domain.OrderDetail{},
		&domain.Pembayaran{},
		&domain.Pengiriman{},
	); err != nil {
		log.Fatalf("[database] auto-migrate failed during model migration: %v", err)
	}

	log.Println("[database] auto-migration (gorm) finished successfully")
}
