package repositories

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/felipepalacio293/stocks-app/models"
	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Create(stock *models.Stock) error {
	return r.db.Create(stock).Error
}

func (r *StockRepository) ListAll() ([]models.Stock, error) {
	var stocks []models.Stock
	if err := r.db.Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *StockRepository) List(page, pageSize int) ([]models.Stock, int64, error) {
	var stocks []models.Stock
	var count int64

	if err := r.db.Model(&models.Stock{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := r.db.Offset(offset).Limit(pageSize).Find(&stocks).Error; err != nil {
		return nil, 0, err
	}

	return stocks, count, nil
}

func (r *StockRepository) Update(stock *models.Stock) error {
	return r.db.Save(stock).Error
}

func (r *StockRepository) BatchInsert(ctx context.Context, stocks []models.Stock, batchSize int) error {
	if len(stocks) == 0 {
		return nil
	}

	if batchSize <= 0 {
		batchSize = 100
	}

	for i := 0; i < len(stocks); i += batchSize {
		end := i + batchSize
		if end > len(stocks) {
			end = len(stocks)
		}

		batch := stocks[i:end]

		err := r.withRetry(ctx, func(tx *gorm.DB) error {
			for _, stock := range batch {
				var existingStock models.Stock
				result := tx.WithContext(ctx).Where("ticker = ? AND brokerage = ?",
					stock.Ticker, stock.Brokerage).First(&existingStock)

				if result.Error != nil {
					if errors.Is(result.Error, gorm.ErrRecordNotFound) {
						if err := tx.WithContext(ctx).Create(&stock).Error; err != nil {
							return err
						}
					} else {
						return result.Error
					}
				} else {
					stock.ID = existingStock.ID
					if err := tx.WithContext(ctx).Save(&stock).Error; err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return fmt.Errorf("error processing batch %d-%d: %w", i, end, err)
		}
	}

	return nil
}

func (r *StockRepository) withRetry(ctx context.Context, operation func(*gorm.DB) error) error {
	var err error
	maxRetries := 5

	for attempt := 0; attempt < maxRetries; attempt++ {
		tx := r.db.WithContext(ctx).Begin()
		if tx.Error != nil {
			return tx.Error
		}

		err = operation(tx)

		if err == nil {
			return tx.Commit().Error
		}

		tx.Rollback()

		if isCockroachTransientError(err) {
			backoffMs := 50 * (1 << uint(attempt))
			jitter := rand.Intn(50)
			sleepTime := time.Duration(backoffMs+jitter) * time.Millisecond

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(sleepTime):
				continue
			}
		}

		return err
	}

	return fmt.Errorf("max retries exceeded: %w", err)
}

func isCockroachTransientError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()

	transientErrors := []string{
		"restart transaction",
		"serialization failure",
		"40001",
		"connection reset by peer",
		"context deadline exceeded",
		"no connection to the server",
		"connection refused",
		"deadline exceeded",
		"read-only transaction",
	}

	for _, pattern := range transientErrors {
		if strings.Contains(errMsg, pattern) {
			return true
		}
	}

	return false
}
