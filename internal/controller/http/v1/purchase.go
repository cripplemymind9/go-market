package v1

import (
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"net/http"
)

type purchaseRoutes struct {
	purchaseService service.Purchase
	validator 		*validator.Validate
}

func newPurchaseRoutes(g *gin.RouterGroup, purchaseService service.Purchase, validator *validator.Validate) {
	r := &purchaseRoutes{
		purchaseService: 	purchaseService,
		validator: 			validator,
	}

	g.POST("/make-purchase", r.makePurchase)
	g.GET("/get-user-purchase", r.getUserPurchases)
	g.GET("/get-product-purchase", r.getProductPurchases)
}

type makePurcahseInput struct {
	UserID 		int `json:"user_id"`
	ProductID 	int `json:"product_id"`
	Quantity 	int `json:"quantity"`
}

func (r *purchaseRoutes) makePurchase(c *gin.Context) {
	var input makePurcahseInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.purchaseService.MakePurchase(c.Request.Context(), types.PurchaseMakePurchaseInput{
		UserID: input.UserID,
		ProductID: input.ProductID,
		Quantity: input.Quantity,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		ID int `json:"id"`
	}

	c.JSON(http.StatusCreated, response{
		ID: id,
	})
}

type getUserPurchasesInput struct {
	UserID int `json:"user_id" validate:"required"`
}

func (r *purchaseRoutes) getUserPurchases(c *gin.Context) {
	var input getUserPurchasesInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := r.purchaseService.GetUserPurchases(c.Request.Context(), input.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Purchases []entity.Purchase
	}

	c.JSON(http.StatusOK, response{
		Purchases: purchases,
	})
}

type getProductPurchasesInput struct {
	ProductID int `json:"product_id" validate:"required"`
}

func (r *purchaseRoutes) getProductPurchases(c *gin.Context) {
	var input getProductPurchasesInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	purchases, err := r.purchaseService.GetProductPurchases(c.Request.Context(), input.ProductID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Purchases []entity.Purchase
	}

	c.JSON(http.StatusOK, response{
		Purchases: purchases,
	})
}