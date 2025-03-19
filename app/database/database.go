package database

import (
	"fmt"
	"log/slog"

	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:generate mockery --name=Database --output=./mocks --filename=mock_database.go
type Database interface {
	Model(value any) *gorm.DB
	Where(query any, args ...any) *gorm.DB
	Count(count *int64) *gorm.DB
	Limit(limit int) *gorm.DB
	Offset(offset int) *gorm.DB
	Find(dest any, conds ...any) *gorm.DB
	First(dest any, conds ...any) *gorm.DB
	Create(value any) *gorm.DB
	Save(value any) *gorm.DB
	Delete(value any, conds ...any) *gorm.DB
}

func ConnectDatabase(cfg *config.Config) (Database, error) {
	const funcName string = "ConnectDatabase"

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
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

	slog.Info("database connected", slog.String("function_name", funcName))

	err = db.AutoMigrate(&models.Song{})
	if err != nil {
		slog.Error("database migration error",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
	}

	slog.Info("database migrated", slog.String("function_name", funcName))

	return db, nil
}
