package main

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"FlashHelper/internal/config"
	"FlashHelper/internal/infra"

	_ "FlashHelper/docs"
)

// @title           FlashHelper API
// @version         1.0
// @description     API documentation cho FlashHelper
// @termsOfService  http://swagger.io/terms/
// @BasePath        /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Nhập JWT theo định dạng: Bearer <token>
func main() {
	// Logger
	infra.InitZerolog()

	// Configuration
	cfg := config.LoadConfig()

	// Initialize application
	router, cleanup, err := InitializeServer(
		context.Background(),
		cfg,
	)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to initialize server")
	}

	defer cleanup()

	// Start HTTP server
	addr := fmt.Sprintf(":%s", cfg.AppPort)

	log.Info().
		Str("addr", addr).
		Msg("server started")

	if err := router.Run(addr); err != nil {
		log.Fatal().
			Err(err).
			Msg("server stopped")
	}
}
