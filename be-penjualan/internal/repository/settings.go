package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"insightflow/be-penjualan/internal/database"
	"insightflow/be-penjualan/internal/domain"
)

var ErrSettingsNotFound = errors.New("settings not found")

// SettingsRepository defines database operations for telegram/anomaly settings.
type SettingsRepository interface {
	GetTelegramConfig(ctx context.Context) (*domain.TelegramConfig, error)
	UpsertTelegramConfig(ctx context.Context, cfg *domain.TelegramConfig) error
	GetAnomalyConfig(ctx context.Context, metricKey string) (*domain.AnomalyConfig, error)
	UpsertAnomalyConfig(ctx context.Context, cfg *domain.AnomalyConfig) error
}

type settingsRepo struct{}

// NewSettingsRepository creates a new SettingsRepository backed by pgxpool.
func NewSettingsRepository() SettingsRepository {
	return &settingsRepo{}
}

func (r *settingsRepo) GetTelegramConfig(ctx context.Context) (*domain.TelegramConfig, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, nama_grup, chat_id, aktif, jam_summary, created_at
		FROM app.telegram_config
		ORDER BY aktif DESC, created_at ASC
		LIMIT 1
	`

	var cfg domain.TelegramConfig
	if err := database.Pool.QueryRow(ctx, q).Scan(
		&cfg.ID, &cfg.NamaGrup, &cfg.ChatID, &cfg.Aktif, &cfg.JamSummary, &cfg.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSettingsNotFound
		}
		return nil, err
	}
	return &cfg, nil
}

func (r *settingsRepo) UpsertTelegramConfig(ctx context.Context, cfg *domain.TelegramConfig) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	if cfg.ID == uuid.Nil {
		const insertQ = `
			INSERT INTO app.telegram_config (nama_grup, chat_id, aktif, jam_summary)
			VALUES ($1, $2, $3, $4)
			RETURNING id, created_at
		`
		return database.Pool.QueryRow(ctx, insertQ,
			cfg.NamaGrup, cfg.ChatID, cfg.Aktif, cfg.JamSummary,
		).Scan(&cfg.ID, &cfg.CreatedAt)
	}

	const updateQ = `
		UPDATE app.telegram_config
		SET nama_grup = $1, chat_id = $2, aktif = $3, jam_summary = $4
		WHERE id = $5
	`
	tag, err := database.Pool.Exec(ctx, updateQ,
		cfg.NamaGrup, cfg.ChatID, cfg.Aktif, cfg.JamSummary, cfg.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		cfg.ID = uuid.Nil
		return r.UpsertTelegramConfig(ctx, cfg)
	}
	return nil
}

func (r *settingsRepo) GetAnomalyConfig(ctx context.Context, metricKey string) (*domain.AnomalyConfig, error) {
	if database.Pool == nil {
		return nil, errors.New("database pool not initialised")
	}

	const q = `
		SELECT id, metric_key, threshold_pct, aktif
		FROM app.anomaly_config
		WHERE metric_key = $1
		LIMIT 1
	`

	var cfg domain.AnomalyConfig
	if err := database.Pool.QueryRow(ctx, q, metricKey).Scan(
		&cfg.ID, &cfg.MetricKey, &cfg.ThresholdPct, &cfg.Aktif,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSettingsNotFound
		}
		return nil, err
	}
	return &cfg, nil
}

func (r *settingsRepo) UpsertAnomalyConfig(ctx context.Context, cfg *domain.AnomalyConfig) error {
	if database.Pool == nil {
		return errors.New("database pool not initialised")
	}

	const q = `
		INSERT INTO app.anomaly_config (metric_key, threshold_pct, aktif)
		VALUES ($1, $2, $3)
		ON CONFLICT (metric_key)
		DO UPDATE SET
			threshold_pct = EXCLUDED.threshold_pct,
			aktif = EXCLUDED.aktif
		RETURNING id
	`

	return database.Pool.QueryRow(ctx, q, cfg.MetricKey, cfg.ThresholdPct, cfg.Aktif).Scan(&cfg.ID)
}
