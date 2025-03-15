package routes

import (
	"github.com/felipepalacio293/stocks-app/config"
	"github.com/felipepalacio293/stocks-app/controllers"
	"github.com/felipepalacio293/stocks-app/repositories"
	"github.com/felipepalacio293/stocks-app/services"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.AllowedOrigins[0])
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	stockRepo := repositories.NewStockRepository(db)
	stockService := services.NewStockService(stockRepo, cfg)
	stockController := controllers.NewStockController(stockService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	api := r.Group("/api/v1")
	{
		public := api.Group("/")
		{
			public.GET("/stocks", stockController.ListStocks)
		}
	}

	return r
}
