// @title SaveSphere API
// @version 1.0
// @description Enterprise Digital Savings and Investment Platform API.
// @host localhost:8080
// @BasePath /
package main

import (
	"log"

	//"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/gin-gonic/gin"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/services"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/application/auth/usecases"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/cloudinary"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/config"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/domain/repositories"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/infrastructure/database"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/interfaces/http/controllers"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/jwt"

	_ "github.com/quantumedgethinkers/SAVINGS-SPHERE/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//userRepository := repositories.NewUserRepository(db)

	jwtService := jwt.New(cfg.JWTSecret)

	_ = jwtService

	cloudinaryService, err := cloudinary.New(
		cfg.CloudName,
		cfg.APIKey,
		cfg.APISecret,
	)
	if err != nil {
		log.Println("Cloudinary not configured yet")
	}

	_ = cloudinaryService

	userRepository := repositories.NewUserRepository(db)
	passwordService := services.NewPasswordService()
	registerUseCase := usecases.NewRegisterUseCase(*userRepository, passwordService)
	authController := controllers.NewAuthController(registerUseCase)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
			"app":    cfg.AppName,
		})
	})

	api := router.Group("/api/v1")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authController.Register)
	}

	log.Println("Server running on :" + cfg.AppPort)
	router.Run(":" + cfg.AppPort)
}
