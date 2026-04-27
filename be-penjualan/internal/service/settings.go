package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"insightflow/be-penjualan/internal/domain"
	"insightflow/be-penjualan/internal/dto"
	"insightflow/be-penjualan/internal/repository"
)

const defaultAnomalyMetricKey = "daily_revenue"
const defaultThresholdPct = 10.0

var (
	ErrInvalidJamSummary = errors.New("jam_summary harus format HH:MM")
	ErrInvalidThreshold  = errors.New("threshold_pct harus di antara 0 dan 100")
)

// SettingsService handles telegram and anomaly settings business logic.
type SettingsService interface {
	GetTelegram(ctx context.Context) (*dto.TelegramConfigResponse, error)
	UpdateTelegram(ctx context.Context, req dto.UpdateTelegramConfigRequest) (*dto.TelegramConfigResponse, error)
}

type settingsService struct {
	repo repository.SettingsRepository
}

// NewSettingsService creates a SettingsService.
func NewSettingsService(repo repository.SettingsRepository) SettingsService {
	return &settingsService{repo: repo}
}

func (s *settingsService) GetTelegram(ctx context.Context) (*dto.TelegramConfigResponse, error) {
	telegramCfg, err := s.repo.GetTelegramConfig(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrSettingsNotFound) {
			telegramCfg = &domain.TelegramConfig{
				NamaGrup:   "Default Group",
				ChatID:     0,
				Aktif:      true,
				JamSummary: "07:00",
			}
			if upsertErr := s.repo.UpsertTelegramConfig(ctx, telegramCfg); upsertErr != nil {
				return nil, upsertErr
			}
		} else {
			return nil, err
		}
	}

	anomalyCfg, err := s.repo.GetAnomalyConfig(ctx, defaultAnomalyMetricKey)
	if err != nil {
		if errors.Is(err, repository.ErrSettingsNotFound) {
			anomalyCfg = &domain.AnomalyConfig{
				MetricKey:    defaultAnomalyMetricKey,
				ThresholdPct: defaultThresholdPct,
				Aktif:        true,
			}
			if upsertErr := s.repo.UpsertAnomalyConfig(ctx, anomalyCfg); upsertErr != nil {
				return nil, upsertErr
			}
		} else {
			return nil, err
		}
	}

	return &dto.TelegramConfigResponse{
		ID:           telegramCfg.ID.String(),
		NamaGrup:     telegramCfg.NamaGrup,
		ChatID:       telegramCfg.ChatID,
		Aktif:        telegramCfg.Aktif,
		JamSummary:   normalizeJamSummary(telegramCfg.JamSummary),
		ThresholdPct: anomalyCfg.ThresholdPct,
	}, nil
}

func (s *settingsService) UpdateTelegram(ctx context.Context, req dto.UpdateTelegramConfigRequest) (*dto.TelegramConfigResponse, error) {
	current, err := s.GetTelegram(ctx)
	if err != nil {
		return nil, err
	}

	next := &domain.TelegramConfig{
		NamaGrup:   current.NamaGrup,
		ChatID:     current.ChatID,
		Aktif:      current.Aktif,
		JamSummary: normalizeJamSummary(current.JamSummary),
	}

	if current.ID != "" {
		if id, parseErr := uuid.Parse(current.ID); parseErr == nil {
			next.ID = id
		}
	}

	threshold := current.ThresholdPct

	if req.NamaGrup != nil {
		nama := strings.TrimSpace(*req.NamaGrup)
		if nama == "" {
			return nil, errors.New("nama_grup tidak boleh kosong")
		}
		next.NamaGrup = nama
	}
	if req.ChatID != nil {
		next.ChatID = *req.ChatID
	}
	if req.Aktif != nil {
		next.Aktif = *req.Aktif
	}
	if req.JamSummary != nil {
		jam := strings.TrimSpace(*req.JamSummary)
		if !isValidJamSummary(jam) {
			return nil, ErrInvalidJamSummary
		}
		next.JamSummary = jam
	}
	if req.ThresholdPct != nil {
		if *req.ThresholdPct < 0 || *req.ThresholdPct > 100 {
			return nil, ErrInvalidThreshold
		}
		threshold = *req.ThresholdPct
	}

	if err := s.repo.UpsertTelegramConfig(ctx, next); err != nil {
		return nil, err
	}

	if err := s.repo.UpsertAnomalyConfig(ctx, &domain.AnomalyConfig{
		MetricKey:    defaultAnomalyMetricKey,
		ThresholdPct: threshold,
		Aktif:        true,
	}); err != nil {
		return nil, err
	}

	return &dto.TelegramConfigResponse{
		ID:           next.ID.String(),
		NamaGrup:     next.NamaGrup,
		ChatID:       next.ChatID,
		Aktif:        next.Aktif,
		JamSummary:   normalizeJamSummary(next.JamSummary),
		ThresholdPct: threshold,
	}, nil
}

func isValidJamSummary(v string) bool {
	_, err := time.Parse("15:04", v)
	return err == nil
}

func normalizeJamSummary(v string) string {
	v = strings.TrimSpace(v)
	if v == "" {
		return "07:00"
	}
	if t, err := time.Parse("15:04:05", v); err == nil {
		return t.Format("15:04")
	}
	if t, err := time.Parse("15:04", v); err == nil {
		return t.Format("15:04")
	}
	return v
}
