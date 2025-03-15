package controllers

import (
	"net/http"
	"strconv"

	"github.com/felipepalacio293/stocks-app/services"
	"github.com/felipepalacio293/stocks-app/utils"
	"github.com/gin-gonic/gin"
)

type StockController struct {
	stockService *services.StockService
}

func NewStockController(stockService *services.StockService) *StockController {
	return &StockController{
		stockService: stockService,
	}
}

func (c *StockController) ListStocks(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	stocks, count, err := c.stockService.ListStocks(page, pageSize)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.PaginatedResponse(stocks, page, pageSize, count, "Stocks retrieved successfully"))
}
