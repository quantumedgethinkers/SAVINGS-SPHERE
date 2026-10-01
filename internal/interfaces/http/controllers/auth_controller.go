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
}

func NewAuthController(
	registerUseCase *usecases.RegisterUseCase,
) *AuthController {
	return &AuthController{
		registerUseCase: registerUseCase,
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
