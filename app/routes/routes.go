package routes

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/mikhail-kirillov/Music-info/config"
	_ "github.com/mikhail-kirillov/Music-info/docs"
	"github.com/mikhail-kirillov/Music-info/handlers"
	"github.com/mikhail-kirillov/Music-info/routes/middlewares"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	const funcName string = "SetupRouter"

	router := gin.New()
	router.Use(middlewares.LoggerMiddleware())
	slog.Debug("middleware has been configured",
		slog.String("function_name", funcName))

	router.GET("/songs", func(ctx *gin.Context) {
		handlers.GetSongs(ctx, db)
	})
	router.GET("/songs/:id/lyrics", func(ctx *gin.Context) {
		handlers.GetSongLyrics(ctx, db)
	})
	router.POST("/songs", func(ctx *gin.Context) {
		handlers.AddSong(ctx, cfg, db)
	})
	router.PUT("/songs/:id", func(ctx *gin.Context) {
		handlers.UpdateSong(ctx, db)
	})
	router.DELETE("/songs/:id", func(ctx *gin.Context) {
		handlers.DeleteSong(ctx, db)
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	slog.Debug("router has been configured",
		slog.String("function_name", funcName))

	return router
}
