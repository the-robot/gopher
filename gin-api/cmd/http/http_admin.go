package http

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// setupAdminRouter is to set up gin router.
func setupAdminRouter() *gin.Engine {
	if viper.GetBool("gin.public.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	// Register routers

	return app
}

// RunAdmin is a method to run gin application for admin server.
func runAdmin(wg *sync.WaitGroup) {
	defer wg.Done()
	host := viper.GetString("gin.admin.host")
	port := viper.GetString("gin.admin.port")

	app := setupAdminRouter()
	_ = app.Run(host + ":" + port)
}
