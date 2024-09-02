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

// signUpInput представляет собой модель данных для запроса на регистрацию.
type signUpInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email 	 string `json:"email" validate:"required,email"`
}

// signUp регистрирует нового пользователя
// @Summary User registration
// @Description Register a new user with username, password, and email
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signUpInput true "User registration input"
// @Success 201 {object} v1.authRoutes.signUp.response
// @Failure 400 {object} ErrorResonse "Invalid request body or validation error"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Router /auth/sign-up [post]
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
			newErrorResponse(c, http.StatusBadRequest, err.Error())
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

// signInInput представляет собой модель данных для запроса на вход.
type signInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// signIn выполняет аутентификацию пользователя
// @Summary User login
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body signInInput true "User login input"
// @Success 201 {object} v1.authRoutes.signIn.response
// @Failure 400 {object} ErrorResonse "Invalid credentials or bad request"
// @Failure 500 {object} ErrorResonse "Internal server error"
// @Router /auth/sign-in [post]
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
		var statusCode int
		var message string

		switch err {
		case serviceerrs.ErrUserNotFound, serviceerrs.ErrInvalidPassword:
			statusCode = http.StatusBadRequest
			message = "Invalid credentials"
		default:
			statusCode = http.StatusInternalServerError
			message = "Internal server error"
		}
		newErrorResponse(c, statusCode, message)
		return
	}

	type response struct {
		Token string `json:"token"`
	}

	c.JSON(http.StatusCreated, response{
		Token: token,
	})
}