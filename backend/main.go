package main

import (
	"chorebutler/infrastructure/db"
	"chorebutler/internal/routers"
	"chorebutler/internal/service"
	"chorebutler/pkg/sanctum"
	"chorebutler/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	// Set the local timezone to UTC for the entire application
	time.Local = time.UTC
}

// @title ChoreButler API
// @version 1.0
// @description This is a sample server for ChoreButler API
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	validate := validator.New()

	cfg, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
		return
	}

	utils.ConfigLogger(cfg)

	// Initialize DB
	database := db.NewDatabaseInstance(cfg)

	// Initialize the token
	cryptoSanctum := sanctum.NewCryptoSanctum(cfg)
	tokenSanctum := sanctum.NewTokenSanctum(cryptoSanctum)
	sanctumToken := sanctum.NewSanctumToken(tokenSanctum, cryptoSanctum, database)

	// Initialize services using the service factory pattern (dependency injection also included repository pattern)
	serviceFactory := service.NewServiceFactory(database, cfg, sanctumToken)
	services := serviceFactory.CreateServices()

	// Initialize API server
	server, err := routers.NewApiServer(
		validate,
		cfg,
		services)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create API server")
		return
	}

	// Setup routers and swagger
	server.SetupRouter()
	server.SetupSwagger(cfg.SwaggerUrl)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start HTTP server in a goroutine
	go func() {
		if err := server.Start(cfg.HTTPServerAddress); err != nil {
			log.Error().Err(err).Msg("Server encountered an error")
			stop <- syscall.SIGTERM // Signal for graceful shutdown
		}
	}()

	log.Info().Msgf("Listening and serving HTTP on %s", cfg.HTTPServerAddress)
	<-stop

}
