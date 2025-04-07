package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/felipepalacio293/stocks-app/models"
	"github.com/google/uuid"
)

type APIClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewAPIClient(baseURL, apiKey string) *APIClient {
	return &APIClient{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type StockResponse struct {
	Status   string      `json:"status"`
	Items    []StockData `json:"items"`
	NextPage string      `json:"next_page"`
}

type StockData struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	Brokerage  string `json:"brokerage"`
	Action     string `json:"action"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Time       string `json:"time"`
}

func (c *APIClient) FetchStocks(ctx context.Context) ([]models.Stock, error) {
	stocks := []models.Stock{}
	nextPage := ""
	baseEndpoint := fmt.Sprintf("%s/production/swechallenge/list", c.baseURL)

	for {
		url := baseEndpoint
		if nextPage != "" {
			url = fmt.Sprintf("%s?next_page=%s", baseEndpoint, nextPage)
		}

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating request: %w", err)
		}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error executing request: %w", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("API returned non-OK status: %d", resp.StatusCode)
		}

		var stockResp StockResponse
		if err := json.NewDecoder(resp.Body).Decode(&stockResp); err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}

		for _, data := range stockResp.Items {
			cleanTargetFrom := strings.ReplaceAll(strings.TrimPrefix(data.TargetFrom, "$"), ",", "")
			targetFrom, err := strconv.ParseFloat(cleanTargetFrom, 64)
			if err != nil {
				log.Printf("Warning: could not parse TargetFrom value '%s' for ticker %s: %v", data.TargetFrom, data.Ticker, err)
				continue
			}

			cleanTargetTo := strings.ReplaceAll(strings.TrimPrefix(data.TargetTo, "$"), ",", "")
			targetTo, err := strconv.ParseFloat(cleanTargetTo, 64)
			if err != nil {
				log.Printf("Warning: could not parse TargetTo value '%s' for ticker %s: %v", data.TargetTo, data.Ticker, err)
				continue
			}

			stock := models.Stock{
				ID:         uuid.New(),
				Ticker:     data.Ticker,
				Company:    data.Company,
				Brokerage:  data.Brokerage,
				Action:     data.Action,
				RatingFrom: data.RatingFrom,
				RatingTo:   data.RatingTo,
				TargetFrom: targetFrom,
				TargetTo:   targetTo,
			}
			stocks = append(stocks, stock)
		}

		if stockResp.NextPage == "" || len(stockResp.Items) == 0 {
			break
		}

		nextPage = stockResp.NextPage

		time.Sleep(100 * time.Millisecond)

		log.Printf("Fetched page with %d items. Next page token: %s", len(stockResp.Items), nextPage)
	}

	log.Printf("Successfully fetched a total of %d stocks", len(stocks))
	return stocks, nil
}
