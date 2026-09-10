package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTAccessExpiry  string `mapstructure:"JWT_ACCESS_EXPIRY"`
	JWTRefreshExpiry string `mapstructure:"JWT_REFRESH_EXPIRY"`

	CloudName string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	APIKey    string `mapstructure:"CLOUDINARY_API_KEY"`
	APISecret string `mapstructure:"CLOUDINARY_API_SECRET"`
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
