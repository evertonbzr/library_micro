package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/evertonbzr/library_micro/cmd/module/user/config"
	"github.com/evertonbzr/library_micro/pkg/queue"
)

func main() {
	config.Load(os.Getenv("ENV"))
	slog.Info("Starting user service...", "env", config.ENV, "port", config.PORT)

	ctx := context.Background()

	if err := queue.ConnectNats(ctx, config.NATS_URI, config.NAME); err != nil {
		log.Fatalf("Error on nats connect %s", config.NATS_URI)
	}
}
