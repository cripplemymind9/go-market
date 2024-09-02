package v1

import (
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/cripplemymind9/go-market/internal/entity"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type productRoutes struct {
	productService 	service.Product
	validator 		*validator.Validate
}

func newProductRoutes(g *gin.RouterGroup, productService service.Product, validator *validator.Validate) {
	r := &productRoutes{
		productService: productService,
		validator: 		validator,
	}

	g.POST("/add-product", r.addProduct)
	g.GET("/get-products", r.getAllProducts)
	g.GET("/get-product/:id", r.getProduct)
	g.PUT("/update-product/:id", r.updateProduct)
	g.DELETE("/delete-product/:id", r.deleteProduct)
}

// addProductInput представляет собой модель данных для добавления продукта.
type addProductInput struct {
	Name 		string 	`json:"name" validate:"required"`
	Description string 	`json:"description" validate:"required"`
	Price 		float64 `json:"price" validate:"required"`
	Quantity 	int		`json:"quantity" validate:"required"`
}

// addProduct добавляет новый продукт в каталог
// @Summary Add a new product
// @Description Add a new product with name, description, price, and quantity
// @Tags products
// @Accept json
// @Produce json
// @Param input body addProductInput true "Product input"
// @Success 201 {object} v1.productRoutes.addProduct.response
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/products/add-product [post]
func (r *productRoutes) addProduct(c *gin.Context) {
	var input addProductInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.productService.AddProduct(c.Request.Context(), types.ProductAddProductInput{
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Quantity: input.Quantity,
	})
	if err != nil {
		if err == serviceerrs.ErrProductAlreadyExists {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
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

// getAllProducts возвращает список всех продуктов
// @Summary Get all products
// @Description Retrieve a list of all available products
// @Tags products
// @Produce json
// @Success 200 {object} v1.productRoutes.getAllProducts.response
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/products/get-products [get]
func (r *productRoutes) getAllProducts(c *gin.Context) {
	var products []entity.Product

	products, err := r.productService.GetAllProducts(c.Request.Context())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	type response struct {
		Products []entity.Product
	}

	c.JSON(http.StatusOK, response{
		Products: products,
	})
}

// getProduct возвращает информацию о продукте по его идентификатору
// @Summary Get product by ID
// @Description Retrieve a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} v1.productRoutes.getProduct.response
// @Failure 400 {object} ErrorResonse "Invalid product ID"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/products/get-product/{id} [get]
func (r *productRoutes) getProduct(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, string(err.Error()))
		return
	}

	product, err := r.productService.GetProductById(c.Request.Context(), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Product entity.Product `json:"product"`
	}

	c.JSON(http.StatusOK, response{
		Product: product,
	})
}

// updateProductInput представляет собой модель данных для обновления продукта.
type updateProductInput struct {
	Name 			string 	`json:"name" validate:"required"`
	Description 	string 	`json:"description" validate:"required"`
	Price 			float64 `json:"price" validate:"required"`
	Quantity 		int 	`json:"quantity" validate:"required"`
}

// updateProduct обновляет информацию о продукте по его идентификатору
// @Summary Update product by ID
// @Description Update product details by ID with new name, description, price, and quantity
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param input body updateProductInput true "Product update input"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/products/update-product/{id} [put]
func (r *productRoutes) updateProduct(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, string(err.Error()))
		return
	}

	var input updateProductInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.productService.UpdateProduct(c.Request.Context(), types.ProductUpdateProductInput{
		ID: id,
		Name: input.Name,
		Description: input.Description,
		Price: input.Price,
		Quantity: input.Quantity,
	}); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
	})
}

// deleteProduct удаляет продукт по его идентификатору
// @Summary Delete product by ID
// @Description Delete a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Success message"
// @Failure 400 {object} ErrorResonse "Invalid product ID"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/products/delete-product/{id} [delete]
func (r *productRoutes) deleteProduct(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, string(err.Error()))
		return
	}

	if err := r.productService.DeleteProduct(c.Request.Context(), id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
	})
}