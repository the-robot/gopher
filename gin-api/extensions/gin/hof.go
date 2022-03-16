package gin

import (
	"gingo/apps/db"
	"gingo/apps/repository"
	"gingo/extensions/error"

	"github.com/gin-gonic/gin"
)

type State struct {
	Context    *gin.Context
	Repository *repository.Repository
}

// RouteWrapper is higher order function to pass DB resource to routes in Gin
func RouteWrapper(resource *db.Resource, route func(s *State) (IResponse, error.IError)) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := repository.NewRepository(resource)
		res, err := route(&State{c, repo})

		if err != nil {
			err.IntoResponse(c)
		} else {
			res.ok(c)
		}
	}
}
