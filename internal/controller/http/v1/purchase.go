package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/internal/service/types"
)

type purchaseRoutes struct {
	purchaseService service.Purchase
	validator       *validator.Validate
}

func newPurchaseRoutes(g *gin.RouterGroup, purchaseService service.Purchase, validator *validator.Validate) {
	r := &purchaseRoutes{
		purchaseService: purchaseService,
		validator:       validator,
	}

	g.POST("/make-purchase", r.makePurchase)
	g.GET("/get-user-purchase/:id", r.getUserPurchases)
	g.GET("/get-product-purchase/:id", r.getProductPurchases)
}

// makePurcahseInput представляет собой модель данных для запроса на покупку продукта.
type makePurcahseInput struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

// makePurchase осуществляет покупку продукта
// @Summary Make a purchase
// @Description Allows a user to purchase a product by specifying user ID, product ID, and quantity
// @Tags purchases
// @Accept json
// @Produce json
// @Param input body makePurcahseInput true "Purchase input data"
// @Success 201 {object} v1.purchaseRoutes.makePurchase.response
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/purchase/make-purchase [post]
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
		UserID:    input.UserID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
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

// getUserPurchases возвращает список покупок пользователя
// @Summary Get user purchases
// @Description Retrieve a list of purchases made by a user specified by user ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} v1.purchaseRoutes.getUserPurchases.response
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/purchase/get-user-purchase/{id} [get]
func (r *purchaseRoutes) getUserPurchases(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, string(err.Error()))
		return
	}

	purchases, err := r.purchaseService.GetUserPurchases(c.Request.Context(), id)
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

// getProductPurchases возвращает список покупок по идентификатору продукта
// @Summary Get product purchases
// @Description Retrieve a list of purchases for a specific product by product ID
// @Tags purchases
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} v1.purchaseRoutes.getProductPurchases.response
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/purchase/get-product-purchase/{id} [get]
func (r *purchaseRoutes) getProductPurchases(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, string(err.Error()))
		return
	}

	purchases, err := r.purchaseService.GetProductPurchases(c.Request.Context(), id)
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
