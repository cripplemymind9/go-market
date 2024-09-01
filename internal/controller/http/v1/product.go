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

type addProductInput struct {
	Name 		string 	`json:"name" validate:"required"`
	Description string 	`json:"description" validate:"required"`
	Price 		float64 `json:"price" validate:"required"`
	Quantity 	int		`json:"quantity" validate:"required"`
}

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

type updateProductInput struct {
	Name 			string 	`json:"name" validate:"required"`
	Description 	string 	`json:"description" validate:"required"`
	Price 			float64 `json:"price" validate:"required"`
	Quantity 		int 	`json:"quantity" validate:"required"`
}

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