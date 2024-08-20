package config

import "github.com/spf13/viper"

type Config struct {
	RedisAddr   string
	JWTSecret   string
	CORSOrigins []string
}

var JWTSecret string

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		RedisAddr:   viper.GetString("REDIS_ADDR"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		CORSOrigins: viper.GetStringSlice("CORS_ORIGINS"),
	}
	
	JWTSecret = cfg.JWTSecret

	return cfg, nil
}
