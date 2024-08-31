package v1

import (
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	userIdCtx = "userId"
)

type AuthMiddleware struct {
	authService service.Auth
}

func (h *AuthMiddleware) UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := bearerToken(c.Request)
		if !ok {
			newErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, err := h.authService.ParseToken(token)
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, "cannot parse token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set(userIdCtx, userId)
		c.Next()
	}
}

func bearerToken(r *http.Request) (string, bool) {
	const prefix = "Bearer "

	header := r.Header.Get("Authorization")
	if header == "" {
		return "", false
	}

	if len(header) > len(prefix) && strings.EqualFold(header[:len(prefix)], prefix) {
		return header[len(prefix):], true
	}

	return "", false
}