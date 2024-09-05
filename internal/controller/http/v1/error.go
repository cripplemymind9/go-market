package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAuthHeader = errors.New("invalid auth header")
	ErrCannotParseToken  = errors.New("cannot parse token")
)

func newErrorResponse(c *gin.Context, errStatus int, message string) {
	err := errors.New(message)

	c.JSON(errStatus, gin.H{"error": err.Error()})
}
