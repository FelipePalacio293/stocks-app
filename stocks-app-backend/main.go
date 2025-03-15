package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/felipepalacio293/stocks-app/clients"
	"github.com/felipepalacio293/stocks-app/config"
	"github.com/felipepalacio293/stocks-app/models"
	"github.com/felipepalacio293/stocks-app/repositories"
	"github.com/felipepalacio293/stocks-app/routes"
	"github.com/felipepalacio293/stocks-app/tasks"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration %v", err)
	}

	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database %v", err)
	}

	err = db.AutoMigrate(&models.Stock{})
	if err != nil {
		log.Fatalf("Failed to migrate database %v", err)
	}

	apiClient := clients.NewAPIClient(
		cfg.APIBaseURL,
		cfg.APIKey,
	)

	stockRepo := repositories.NewStockRepository(db)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	syncTask := tasks.NewStockSyncTask(
		stockRepo,
		apiClient,
		30*time.Minute,
	)
	go syncTask.Start(ctx)
	log.Println("Stock sync task started - will sync every 30 minutes")

	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Running initial stock sync...")
		_, taskCancel := context.WithTimeout(ctx, 60*time.Second)
		defer taskCancel()
		syncTask.SyncStocks(ctx)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		r := routes.SetupRouter(db, cfg)
		log.Printf("Starting server on port %s", cfg.ServerPort)
		err = r.Run(":" + cfg.ServerPort)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
			cancel()
		}
	}()

	<-quit
	log.Println("Shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	cancel()
	log.Println("Background tasks are shutting down...")

	select {
	case <-shutdownCtx.Done():
		log.Println("Shutdown timed out")
	case <-time.After(3 * time.Second):
		log.Println("Graceful shutdown completed")
	}
}
