package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/mikhail-kirillov/Music-info/config"
	"github.com/mikhail-kirillov/Music-info/database"
	"github.com/mikhail-kirillov/Music-info/routes"
)

//	@title			Music info
//	@version		0.0.1
//	@description	Online music library
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Kirillov Mikhail
//	@contact.url	https://mikhail-kirillov.github.io/
//	@contact.email	kirillov.mikhail.job@icloud.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	const funcName string = "main"

	config.InitLogger()
	slog.Debug("logger initialization completed",
		slog.String("function_name", funcName))

	cfg := config.LoadConfig()
	slog.Debug("configuration loaded",
		slog.String("function_name", funcName))

	db, err := database.ConnectDatabase(cfg)
	if err != nil {
		slog.Error("database connection error",
			slog.String("function_name", funcName))
		log.Fatal("connection to database error:", err)
	}
	slog.Debug("database is connected",
		slog.String("function_name", funcName))

	router := routes.SetupRouter(cfg, db)
	slog.Debug("router has been configured",
		slog.String("function_name", funcName))

	address := fmt.Sprintf("0.0.0.0:%s", cfg.ServerPort)
	slog.Info("server started",
		slog.String("function_name", funcName))
	err = router.Run(address)
	if err != nil {
		slog.Error("run error",
			slog.String("function_name", funcName),
			slog.String("error", err.Error()))
	}
	slog.Info("server stopped",
		slog.String("function_name", funcName))
}
