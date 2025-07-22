package router

import (
	"golang-subscriptions-api/internal/handler"

	_ "golang-subscriptions-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter sets up the Gin router with all routes
func SetupRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	subscriptions := r.Group("/subscriptions")
	{
		subscriptions.POST("", h.Create)
		subscriptions.GET("", h.List)
		subscriptions.GET("/:id", h.Get)
		subscriptions.PUT("/:id", h.Update)
		subscriptions.DELETE("/:id", h.Delete)
	}

	r.GET("/subscriptions/total", h.TotalSum)

	return r
}
