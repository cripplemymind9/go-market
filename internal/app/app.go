package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/cripplemymind9/go-market/config"
	v1 "github.com/cripplemymind9/go-market/internal/controller/http/v1"
	"github.com/cripplemymind9/go-market/internal/repository"
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/pkg/hasher"
	"github.com/cripplemymind9/go-market/pkg/httpserver"
	"github.com/cripplemymind9/go-market/pkg/postgres"
)

func Run(configPath string) {
	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize config")
	}

	// logger
	SetLogrus(cfg.Log.Level)

	// Postgres
	log.Info("Initializing PostgreSQL...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.WithError(fmt.Errorf("app - Run - postgres.New: %w", err)).Fatal("Failed to initialize PostgreSQL")
	}
	defer pg.Close()

	//  Migrations
	err = InitMigrations()
	if err != nil {
		log.WithError(err).Fatal("Failed to initialize migrations")
	}

	// Repositories
	log.Info("Initializing repositories...")
	repositories := repository.NewRepositories(pg)

	// Services dependencies
	deps := service.ServiceDependencies{
		Repos:    *repositories,
		Hasher:   hasher.NewBcryptHasher(),
		SignKey:  cfg.JWT.SignKey,
		TokenTTL: cfg.JWT.TokenTTL,
	}
	services := service.NewServices(deps)

	// Validator
	validator := validator.New()

	// Gin router
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("validator", validator)
		c.Next()
	})
	v1.NewRouter(router, services, validator)

	// HTTP server
	log.Infof("Starting HTTP server on port %s...", cfg.HTTP.Port)
	httpServer := httpserver.New(router, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Infof("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Infof("app - Run - httpServer.Notify: %v", err)
	}

	// Graceful shutdown
	log.Info("Shutting down HTTP server...")
	if err := httpServer.Shutdown(); err != nil {
		log.WithError(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err)).Error("Failed to shutdown HTTP server")
	}
}
