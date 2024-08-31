package v1

import (
	"github.com/cripplemymind9/go-market/internal/service"
	"github.com/cripplemymind9/go-market/internal/service/serviceerrs"
	"github.com/cripplemymind9/go-market/internal/service/types"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authRoutes struct {
	authService service.Auth
	validator 	*validator.Validate
}

func newAuthRoutes(g *gin.RouterGroup, authService service.Auth, validator *validator.Validate) {
	r := &authRoutes{
		authService: 	authService,
		validator: 		validator,
	}

	g.POST("/sign-up", r.signUp)
	g.POST("/sign-in", r.signIn)
}

type signUpInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"reauired"`
	Email 	 string `json:"email" validate:"required,email"`
}

func (r *authRoutes) signUp(c *gin.Context) {
	var input signUpInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.authService.RegisterUser(c.Request.Context(), types.AuthRegisterUserInput{
		Username: 	input.Username,
		Password: 	input.Password,
		Email: 		input.Email,
	})
	if err != nil {
		if err == serviceerrs.ErrUserAlreadyExists {
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

type signInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *authRoutes) signIn(c *gin.Context) {
	var input signInInput

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := r.validator.Struct(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := r.authService.GenerateToken(c.Request.Context(), types.AuthGenerateTokenInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		if err == serviceerrs.ErrUserNotFound {
			newErrorResponse(c, http.StatusBadRequest, "invalid username or password")
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	c.JSON(http.StatusCreated, response{
		Token: token,
	})
}