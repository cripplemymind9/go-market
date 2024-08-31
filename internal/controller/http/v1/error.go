package v1

import (
	"github.com/gin-gonic/gin"
	"errors"
)
var (
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrCannotParseToken  = errors.New("cannot parse token")
)

func newErrorResponse(c *gin.Context, errStatus int, message string) {
	c.JSON(errStatus, gin.H{"error": message})
}