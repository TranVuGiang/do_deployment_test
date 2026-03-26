package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	db    *pgxpool.Pool
	redis *redis.Client
}

type ReadinessCheckResponse struct {
	Status      string      `json:"status"`
	Dependecies Dependecies `json:"dependencies"`
}

type Dependecies struct {
	Postgres string `json:"postgres"`
	Valkey   string `json:"valkey"`
}

func NewService(db *pgxpool.Pool, redis *redis.Client) *Service {
	return &Service{
		db: db, redis: redis,
	}
}

func (h *Service) CheckReadiness(ctx context.Context) *ReadinessCheckResponse {
	deps := Dependecies{
		Postgres: "ok",
		Valkey:   "ok",
	}

	overall := "ready"

	if err := h.db.Ping(ctx); err != nil {
		deps.Postgres = "failed"
		overall = "not_ready"
	}

	if err := h.redis.Ping(ctx).Err(); err != nil {
		deps.Valkey = "failed"
		overall = "not_ready"
	}

	return &ReadinessCheckResponse{Status: overall, Dependecies: deps}
}
