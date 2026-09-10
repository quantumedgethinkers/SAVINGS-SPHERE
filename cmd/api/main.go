package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/cloudinary"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/config"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/infrastructure/database"
	"github.com/quantumedgethinkers/SAVINGS-SPHERE/internal/jwt"
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

	_ = db

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

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
			"app":    cfg.AppName,
		})
	})

	log.Println("Server running on :" + cfg.AppPort)
	router.Run(":" + cfg.AppPort)
}
