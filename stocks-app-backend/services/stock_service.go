package services

import (
	"math"
	"sort"
	"time"

	"github.com/felipepalacio293/stocks-app/config"
	"github.com/felipepalacio293/stocks-app/models"
	"github.com/felipepalacio293/stocks-app/repositories"
)

type StockRecommendation struct {
	Ticker        string  `json:"ticker"`
	Company       string  `json:"company"`
	Score         float64 `json:"score"`
	Rating        string  `json:"rating"`
	TargetPrice   float64 `json:"target_price"`
	Action        string  `json:"action"`
	ChangePercent float64 `json:"change_percent"`
}

type StockService struct {
	repo *repositories.StockRepository
	cfg  *config.Config
}

const (
	// Factores de la puntuación
	TargetChangePercentMultiplier float64 = 2.0
	TargetPriceLogMultiplier      float64 = 2.0

	// Puntuaciones para las acciones
	ActionUpgradedScore      float64 = 5.0
	ActionTargetRaisedScore  float64 = 3.0
	ActionReiteratedScore    float64 = 1.0
	ActionTargetLoweredScore float64 = -2.0
	ActionDowngradedScore    float64 = -4.0

	// Puntuaciones por tiempo
	RecentUpdateDaysThreshold           int     = 7
	RecentUpdateScore                   float64 = 2.0
	ModeratelyRecentUpdateDaysThreshold int     = 30
	ModeratelyRecentUpdateScore         float64 = 1.0
)

var RatingScores = map[string]float64{
	"Buy":            5.0,
	"Outperform":     4.0,
	"Overweight":     4.0,
	"Equal Weight":   3.0,
	"Sector Perform": 3.0,
	"Hold":           2.0,
	"Underperform":   1.0,
	"Sell":           0.0,
}

func NewStockService(repo *repositories.StockRepository, cfg *config.Config) *StockService {
	return &StockService{
		repo: repo,
		cfg:  cfg,
	}
}

func (s *StockService) ListStocks(page int, pageSize int, ticker string) ([]models.StockResponse, int64, error) {
	stocks, count, err := s.repo.List(page, pageSize, ticker)

	if err != nil {
		return nil, 0, err
	}

	stockResponses := make([]models.StockResponse, len(stocks))
	for i, stock := range stocks {
		stockResponses[i] = stock.ToResponse()
	}

	return stockResponses, count, nil
}

func (s *StockService) GetAllStocks() ([]models.Stock, error) {
	return s.repo.ListAll()
}

func (s *StockService) GetStockRecommendations(topN int) ([]StockRecommendation, error) {
	stocks, err := s.repo.ListAll()
	if err != nil {
		return nil, err
	}

	return recommendStocks(stocks, topN), nil
}

func recommendStocks(stocks []models.Stock, topN int) []StockRecommendation {
	scoredStocks := make([]StockRecommendation, 0, len(stocks))

	for _, stock := range stocks {
		score := scoreStock(stock)
		scoredStocks = append(scoredStocks, score)
	}

	sort.Slice(scoredStocks, func(i, j int) bool {
		return scoredStocks[i].Score > scoredStocks[j].Score
	})

	if len(scoredStocks) < topN {
		topN = len(scoredStocks)
	}

	return scoredStocks[:topN]
}

func (s *StockService) GetTopStocksByAction(stocks []models.Stock, action string, topN int) []StockRecommendation {
	filtered := make([]models.Stock, 0)

	for _, stock := range stocks {
		if stock.Action == action {
			filtered = append(filtered, stock)
		}
	}

	return recommendStocks(filtered, topN)
}

func (s *StockService) GetTopStocksByBrokerage(stocks []models.Stock, brokerage string, topN int) []StockRecommendation {
	filtered := make([]models.Stock, 0)

	for _, stock := range stocks {
		if stock.Brokerage == brokerage {
			filtered = append(filtered, stock)
		}
	}

	return recommendStocks(filtered, topN)
}

func (s *StockService) GetTopStocksByRating(stocks []models.Stock, minRating string, topN int) []StockRecommendation {
	filtered := make([]models.Stock, 0)
	ratingValues := map[string]int{
		"Buy":            5,
		"Outperform":     4,
		"Overweight":     4,
		"Equal Weight":   3,
		"Sector Perform": 3,
		"Hold":           2,
		"Underperform":   1,
		"Sell":           0,
	}

	minRatingValue, exists := ratingValues[minRating]
	if !exists {
		minRatingValue = 3
	}

	for _, stock := range stocks {
		stockRatingValue, exists := ratingValues[stock.RatingTo]
		if exists && stockRatingValue >= minRatingValue {
			filtered = append(filtered, stock)
		}
	}

	return recommendStocks(filtered, topN)
}

func scoreStock(stock models.Stock) StockRecommendation {
	var score float64 = 0

	targetChange := stock.TargetTo - stock.TargetFrom
	var targetChangePercent float64 = 0
	if stock.TargetFrom > 0 {
		targetChangePercent = (targetChange / stock.TargetFrom) * 100
		score += targetChangePercent * TargetChangePercentMultiplier
	}

	if ratingScore, exists := RatingScores[stock.RatingTo]; exists {
		score += ratingScore
	}

	switch stock.Action {
	case "target raised by":
		score += 3
	case "upgraded by":
		score += 5
	case "reiterated by":
		score += 1
	case "target lowered by":
		score -= 2
	case "downgraded by":
		score -= 4
	}

	if stock.TargetTo > 0 {
		relativeTargetScore := math.Log(stock.TargetTo) * TargetPriceLogMultiplier
		score += relativeTargetScore
	}

	updateTime := stock.UpdatedAt
	now := time.Now()
	daysDifference := int(now.Sub(updateTime).Hours() / 24)

	if daysDifference <= RecentUpdateDaysThreshold {
		score += 2
	} else if daysDifference <= ModeratelyRecentUpdateDaysThreshold {
		score += 1
	}

	return StockRecommendation{
		Ticker:        stock.Ticker,
		Company:       stock.Company,
		Score:         score,
		Rating:        stock.RatingTo,
		TargetPrice:   stock.TargetTo,
		Action:        stock.Action,
		ChangePercent: targetChangePercent,
	}
}
