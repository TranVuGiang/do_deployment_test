package main

import (
	"context"
	"log"

	"github.com/TranVuGiang/digital_project_deploy/internal/config"
	"github.com/TranVuGiang/digital_project_deploy/internal/handler"
	"github.com/TranVuGiang/digital_project_deploy/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.New(nil)
	if err != nil {
		log.Fatalf("failed to parse env config: %v", err)
	}

	ctx := context.Background()
	db, err := pgxpool.New(ctx, cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect with postgres: %v", err)
	}
	defer db.Close()

	valkey := redis.NewClient(&redis.Options{
		Addr:     cfg.ValkeyAddr,
		Password: cfg.ValkeyPassword,
	})
	defer valkey.Close()

	svc := service.NewService(db, valkey)

	r := gin.Default()

	r.GET("/health", handler.HealthCheck)
	r.GET("/readiness", handler.ReadinessHandler(svc))

	addrs := ":" + cfg.Port
	log.Printf("server starting on: %s", addrs)
	if err := r.Run(addrs); err != nil {
		log.Fatalf("server error connect: %v", err)
	}
}
