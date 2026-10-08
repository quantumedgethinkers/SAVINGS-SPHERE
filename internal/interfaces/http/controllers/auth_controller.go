package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/dto"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/usecases"
)

type AuthController struct {
	registerUseCase *usecases.RegisterUseCase
	loginUseCase    *usecases.LoginUseCase
}

func NewAuthController(
	registerUseCase *usecases.RegisterUseCase,
	loginUsecase *usecases.LoginUseCase,
) *AuthController {
	return &AuthController{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUsecase,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new SaveSphere user account.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration details"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := c.registerUseCase.Execute(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		if errors.Is(err, usecases.ErrEmailAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "email already exists",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Login user
// @Description Authenticates a user and returns an access token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/v1/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := c.loginUseCase.Execute(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		if errors.Is(err, usecases.ErrInvalidCredentials) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid email or password",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
