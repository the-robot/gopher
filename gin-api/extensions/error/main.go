package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IError interface {
	Error() string
	IntoResponse(c *gin.Context)
}

func Internal(message string, err error) IError {
	return newError(internal, message, err)
}

func NotFound(message string, err error) IError {
	return newError(notFound, message, err)
}

func InvalidData(message string, err error) IError {
	return newError(invalidData, message, err)
}

func InvalidCredential(message string, err error) IError {
	return newError(invalidCredential, message, err)
}

type err struct {
	errorCode  code
	debugError error
	Message    string `json:"message"`
}

func (e *err) getHttpStatus() int {
	switch e.errorCode {
	case internal:
		return http.StatusInternalServerError
	case notFound:
		return http.StatusNotFound
	default:
		return http.StatusBadRequest
	}
}

func (e *err) Error() string {
	return fmt.Sprintf("[ %d ]: %s", e.errorCode, e.Message)
}

func (e *err) IntoResponse(c *gin.Context) {
	res := gin.H{
		"status":  "error",
		"message": e.Message,
	}

	// in non-production environment, also return actual error information for debugging purpose
	if viper.GetBool("gin.client.debug") {
		if e.debugError != nil {
			res["error"] = e.debugError.Error()
		} else {
			res["error"] = nil
		}
	}

	c.JSON(e.getHttpStatus(), res)
}
