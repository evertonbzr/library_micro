package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/evertonbzr/library_micro/cmd/module/user/config"
	"github.com/evertonbzr/library_micro/internal/user/subscriber"
	"github.com/evertonbzr/library_micro/pkg/infra/db"
	"github.com/evertonbzr/library_micro/pkg/infra/queue"
	"github.com/evertonbzr/library_micro/pkg/infra/redis"
)

func main() {
	config.Load(os.Getenv("ENV"))
	slog.Info("Starting user service...", "env", config.ENV, "port", config.PORT)

	ctx := context.Background()

	// Connect to nats
	if err := queue.ConnectNats(ctx, config.NATS_URI, config.NAME); err != nil {
		log.Fatalf("Error on nats connect %s", config.NATS_URI)
	}
	slog.Info("Nats connected", "uri", config.NATS_URI)

	// Connect to redis
	if _, err := redis.ConnectRedisClient(config.REDIS_URL); err != nil {
		log.Fatalf("Error on redis connect %s", config.REDIS_URL)
	}
	slog.Info("Redis connected", "uri", config.REDIS_URL)

	// Connect to database
	if _, err := db.New(config.DATABASE_URL, config.IsDevelopment()); err != nil {
		log.Fatalf("Error on database connect %s", config.DATABASE_URL)
	}
	slog.Info("Database connected PostgreSQL")

	if config.IsDevelopment() {
		db.Migrate()
	}

	queue.ListenSubscriber(subscriber.GetAll()...)

	c := make(chan os.Signal, 1)

	<-c
}
