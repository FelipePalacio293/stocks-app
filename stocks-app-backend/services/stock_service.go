package services

import (
	"github.com/felipepalacio293/stocks-app/config"
	"github.com/felipepalacio293/stocks-app/models"
	"github.com/felipepalacio293/stocks-app/repositories"
)

type StockService struct {
	repo *repositories.StockRepository
	cfg  *config.Config
}

func NewStockService(repo *repositories.StockRepository, cfg *config.Config) *StockService {
	return &StockService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *StockService) ListStocks(page, pageSize int) ([]models.StockResponse, int64, error) {
	stocks, count, err := s.repo.List(page, pageSize)

	if err != nil {
		return nil, 0, err
	}

	stockResponses := make([]models.StockResponse, len(stocks))
	for i, stock := range stocks {
		stockResponses[i] = stock.ToResponse()
	}

	return stockResponses, count, nil
}
