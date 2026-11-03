package router

import (
	"github.com/gin-gonic/gin"

	"FlashHelper/internal/delivery/api/handler"
)

func SetupUserRoutes(
	r *gin.RouterGroup,
	userHandler *handler.UserHandler,
) {
	userRoutes := r.Group("/users")

	{
		userRoutes.POST("/", userHandler.Create)
		userRoutes.GET("/", userHandler.List)
		userRoutes.GET("/:id", userHandler.GetByID)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
	}
}