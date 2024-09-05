package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/cripplemymind9/go-market/docs"
	"github.com/cripplemymind9/go-market/internal/service"
)

type ErrorResonse struct {
	Error string `json:"error"`
}

func NewRouter(router *gin.Engine, services *service.Services, validator *validator.Validate) {
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		newAuthRoutes(auth, services.Auth, validator)
	}

	authMiddleware := &AuthMiddleware{services.Auth}
	v1 := router.Group("/api/v1", authMiddleware.UserIdentity())
	{
		newProductRoutes(v1.Group("/products"), services.Product, validator)
		newPurchaseRoutes(v1.Group("/purchase"), services.Purchase, validator)
	}
}
