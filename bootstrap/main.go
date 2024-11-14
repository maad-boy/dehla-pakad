package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/maad-boy/dehla-pakad/internal/router"
)

func RunApp() {
	app := gin.Default()
	router.AddRouter(app)
	app.Run(":8080")
}
