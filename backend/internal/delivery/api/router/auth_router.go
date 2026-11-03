package router

import (
	"github.com/gin-gonic/gin"

	"FlashHelper/internal/delivery/api/handler"
)

func SetupAuthRoutes(
	r *gin.RouterGroup,
	authHandler *handler.AuthHandler,
	authMiddleware gin.HandlerFunc,
) {
	authRoutes := r.Group("/auth")

	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/logout", authMiddleware, authHandler.Logout)
	}
}
