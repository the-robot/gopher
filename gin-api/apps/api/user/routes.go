package user

import (
	"gingo/extensions/error"
	"gingo/extensions/gin"
)

func getHelloWorld(s *gin.State) (gin.IResponse, error.IError) {
	return gin.Ok("Hello World"), nil
}
