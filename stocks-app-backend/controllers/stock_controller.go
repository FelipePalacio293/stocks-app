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

func (c *StockController) GetStockRecommendations(ctx *gin.Context) {
	topNStr := ctx.DefaultQuery("top_n", "5")

	topN, err := strconv.Atoi(topNStr)
	if err != nil || topN < 1 {
		topN = 5
	}

	stockRecommendations, err := c.stockService.GetStockRecommendations(topN)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.SuccessResponse(stockRecommendations, "Stock recommendations retrieved successfully"))
}

func (c *StockController) GetRecommendationsByAction(ctx *gin.Context) {
	action := ctx.Param("action")
	topNStr := ctx.DefaultQuery("limit", "3")

	topN, err := strconv.Atoi(topNStr)
	if err != nil || topN <= 0 {
		topN = 3
	}

	stocks, err := c.stockService.GetAllStocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	recommendations := c.stockService.GetTopStocksByAction(stocks, action, topN)

	ctx.JSON(http.StatusOK, utils.SuccessResponse(recommendations, "Stock recommendations by action retrieved successfully"))
}

func (c *StockController) GetRecommendationsByBrokerage(ctx *gin.Context) {
	brokerage := ctx.Param("brokerage")
	topNStr := ctx.DefaultQuery("limit", "3")

	topN, err := strconv.Atoi(topNStr)
	if err != nil || topN <= 0 {
		topN = 3
	}

	stocks, err := c.stockService.GetAllStocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	recommendations := c.stockService.GetTopStocksByBrokerage(stocks, brokerage, topN)

	ctx.JSON(http.StatusOK, utils.SuccessResponse(recommendations, "Stock recommendations by brokerage retrieved successfully"))
}

func (c *StockController) GetRecommendationsByRating(ctx *gin.Context) {
	rating := ctx.Param("rating")
	topNStr := ctx.DefaultQuery("limit", "3")

	topN, err := strconv.Atoi(topNStr)
	if err != nil || topN <= 0 {
		topN = 3
	}

	stocks, err := c.stockService.GetAllStocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err.Error()))
		return
	}

	recommendations := c.stockService.GetTopStocksByRating(stocks, rating, topN)

	ctx.JSON(http.StatusOK, utils.SuccessResponse(recommendations, "Stock recommendations by rating retrieved successfully"))
}
