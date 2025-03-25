package database

import (
	"fmt"
	"log/slog"

	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	const funcName string = "ConnectDatabase"

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBPort,
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		slog.Error("error connecting to database",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
		return nil, err
	}

	slog.Info("database connected",
		slog.String("function_name", funcName))

	err = db.AutoMigrate(&models.SongTable{})
	if err != nil {
		slog.Error("database migration error",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
	}

	slog.Info("database migrated",
		slog.String("function_name", funcName))

	return db, nil
}
