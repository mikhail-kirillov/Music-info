package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBName      string
	DBPassword  string
	ServerPort  string
	MusicApiURL string
}

func LoadConfig() *Config {
	const funcName string = "LoadConfig"
	viper.AutomaticEnv()
	slog.Debug("returned config", slog.String("function_name", funcName))
	return &Config{
		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("POSTGRES_PORT"),
		DBUser:      viper.GetString("POSTGRES_USER"),
		DBName:      viper.GetString("POSTGRES_DB"),
		DBPassword:  viper.GetString("POSTGRES_PASSWORD"),
		ServerPort:  viper.GetString("SERVER_PORT"),
		MusicApiURL: viper.GetString("MUSIC_API_URL"),
	}
}

func InitLogger() {
	const funcName string = "InitLogger"
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Debug("logger has been set", slog.String("function_name", funcName))
}
