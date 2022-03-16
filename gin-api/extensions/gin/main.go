package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResponse interface {
	ok(c *gin.Context)
}

func Ok(data interface{}) IResponse {
	return &response{httpStatus: http.StatusOK, data: data}
}

func Created(data interface{}) IResponse {
	return &response{httpStatus: http.StatusCreated, data: data}
}

func Accepted(data interface{}) IResponse {
	return &response{httpStatus: http.StatusAccepted, data: data}
}

func NoContent(data interface{}) IResponse {
	return &response{httpStatus: http.StatusNoContent, data: data}
}

type response struct {
	httpStatus int
	data       interface{}
}

func (r *response) ok(c *gin.Context) {
	c.JSON(r.httpStatus, gin.H{
		"status": "ok",
		"data":   r.data,
	})
}
