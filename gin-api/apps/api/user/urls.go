package user

import (
	"gingo/apps/db"
	ginExt "gingo/extensions/gin"

	"github.com/gin-gonic/gin"
)

// Register api routes.
func Register(router *gin.RouterGroup, resource *db.Resource) {
	router.GET("/", ginExt.RouteWrapper(resource, getHelloWorld))
}
