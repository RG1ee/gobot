package configs

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Load() error {
	if err := loadEnv(); err != nil {
		return err
	}

	return nil
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	viper.AutomaticEnv()
	return nil
}
