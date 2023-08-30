package main

import (
	"github.com/0B1t322/items/internal/controllers"
	"github.com/0B1t322/items/internal/gorm/sqlite"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func main() {
	logger := slog.Default()
	db, err := sqlite.ConnectToDatabase("db/gorm.db")
	if err != nil {
		panic(err)
	}

	if err := sqlite.ApplyMigrations(db); err != nil {
		panic(err)
	}

	app := gin.Default()

	itemsController := controllers.NewItemController(db, logger)

	api := app.Group("api")
	{
		itemsController.Build(api)
	}

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
