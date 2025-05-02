package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{
		"error":   true,
		"message": message,
	})
}
func BadRequest(c *gin.Context, message string) {
	HttpError(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context, message string) {
	HttpError(c, http.StatusNotFound, message)
}

func InternalServerError(c *gin.Context, message string) {
	HttpError(c, http.StatusInternalServerError, message)
}
