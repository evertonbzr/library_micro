package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/evertonbzr/library_micro/internal/user/api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/valyala/fasthttp"
)

type APIConfig struct {
	Port string
}

func Start(cfg *APIConfig) {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(healthcheck.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	authHandler := handler.NewAuthHandler()
	userHandler := handler.NewUserHandler()

	app.Post("/sign-in", authHandler.SignIn())
	app.Post("/sign-up", authHandler.SignUp())
	app.Get("/me", userHandler.GetMe())

	// Server setup
	srv := &fasthttp.Server{
		Handler: app.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			log.Fatalf("Error on listen and serve %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := srv.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	default:
		log.Println("Server exiting")
	}
}
