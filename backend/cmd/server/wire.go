//go:build wireinject

//go:generate wire

package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"FlashHelper/internal/config"
	"FlashHelper/internal/delivery/api/handler"
	"FlashHelper/internal/delivery/api/router"
	"FlashHelper/internal/infra/postgres"
	"FlashHelper/internal/usecase/implement"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ProvideRouter(
	userHandler *handler.UserHandler,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept", "Authorization",
		},
		AllowCredentials: true,
	}))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	router.SetupUserRoutes(api, userHandler)

	return r
}

func InitializeServer(
	ctx context.Context,
	cfg *config.Config,
) (*gin.Engine, func(), error) {
	wire.Build(
		// Infrastructure
		postgres.NewDB,

		// Repository
		postgres.NewUserRepository,

		// UseCase
		implement.NewUserService,

		// Handler
		handler.NewUserHandler,

		// Router
		ProvideRouter,
	)

	return nil, nil, nil
}
