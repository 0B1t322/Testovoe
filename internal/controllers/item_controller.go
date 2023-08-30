package controllers

import (
	"github.com/0B1t322/items/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strings"
)

type ItemController struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewItemController(db *gorm.DB, logger *slog.Logger) *ItemController {
	return &ItemController{
		db:     db,
		logger: logger.With("controller", "ItemsController"),
	}
}

func (i ItemController) Build(r gin.IRouter) {
	r.GET("get-items", i.GetItems)
}

func (i ItemController) GetItems(c *gin.Context) {
	ids := strings.Split(c.Query("id"), ",")

	var items []models.Item
	db := i.db
	{
		if len(ids) > 0 {
			db = db.Where(`id in ?`, ids)
		}

		if err := db.Find(&items).Error; err != nil {
			i.logger.Error("Failed to get items", "err", err)
			c.String(http.StatusInternalServerError, "Failed to get items")
			c.Abort()
			return
		}
	}

	c.JSON(http.StatusOK, items)
}
