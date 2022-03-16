package http

import (
	"sync"

	"gingo/extensions/authentication"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	userAPI "gingo/apps/api/user"
	"gingo/apps/db"
)

// setupPublicRouter is to set up gin router.
func setupPublicRouter() *gin.Engine {
	if viper.GetBool("gin.public.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	// set authentication handler
	secret := viper.GetString("gin.public.jwt.secret")
	noAuthRoutes := viper.GetStringSlice("gin.public.noauth.routes")
	authenticator := authentication.New(secret, noAuthRoutes)
	app.Use(authenticator.ValidateToken)

	// setup database connection
	resource, err := db.Init()
	if err != nil {
		log.Error(err)
	}
	defer resource.Close()

	// authentication router
	authRouter := app.Group("/auth")
	authenticator.Register(authRouter, resource)

	// user routers
	userRouter := app.Group("/user")
	userAPI.Register(userRouter, resource)

	return app
}

// RunPublic is a method to run gin application for user server.
func runPublic(wg *sync.WaitGroup) {
	defer wg.Done()
	host := viper.GetString("gin.public.host")
	port := viper.GetString("gin.public.port")

	app := setupPublicRouter()
	_ = app.Run(host + ":" + port)
}
