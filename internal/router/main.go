package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maad-boy/dehla-pakad/internal/router/handlers"
	"github.com/maad-boy/dehla-pakad/internal/router/handlers/health"
)

func AddRouter(engin *gin.Engine) {
	engin.GET("/health", handlers.GinHandler(health.NewHandler()))
}