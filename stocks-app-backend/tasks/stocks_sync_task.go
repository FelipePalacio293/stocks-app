package tasks

import (
	"context"
	"log"
	"time"

	"github.com/felipepalacio293/stocks-app/clients"
	"github.com/felipepalacio293/stocks-app/repositories"
)

type StockSyncTask struct {
	stockRepo *repositories.StockRepository
	apiClient *clients.APIClient
	interval  time.Duration
}

func NewStockSyncTask(stockRepo *repositories.StockRepository, apiClient *clients.APIClient, interval time.Duration) *StockSyncTask {
	return &StockSyncTask{
		stockRepo: stockRepo,
		apiClient: apiClient,
		interval:  interval,
	}
}

func (t *StockSyncTask) Start(ctx context.Context) {
	t.SyncStocks(ctx)

	ticker := time.NewTicker(t.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			t.SyncStocks(ctx)
		case <-ctx.Done():
			log.Println("Stock sync task stopped")
			return
		}
	}
}

func (t *StockSyncTask) SyncStocks(ctx context.Context) {
	log.Println("Syncing stocks from API...")

	stocks, err := t.apiClient.FetchStocks(ctx)
	if err != nil {
		log.Printf("Error fetching stocks: %v", err)
		return
	}

	err = t.stockRepo.BatchInsert(ctx, stocks, 100)
	if err != nil {
		log.Printf("Error storing stocks: %v", err)
		return
	}

	log.Printf("Successfully synced %d stocks", len(stocks))
}
