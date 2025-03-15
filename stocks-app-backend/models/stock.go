package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Ticker    string `json:"ticker" gorm:"size:20;index:idx_stock_ticker"`
	Company   string `json:"company" gorm:"size:255;index:idx_stock_company"`
	Brokerage string `json:"brokerage" gorm:"size:100;column:brokerage"`

	Action     string `json:"action" gorm:"size:50"`
	RatingFrom string `json:"rating_from" gorm:"size:50"`
	RatingTo   string `json:"rating_to" gorm:"size:50"`

	TargetFrom float64 `json:"target_from" gorm:"type:decimal(10,2)"`
	TargetTo   float64 `json:"target_to" gorm:"type:decimal(10,2)"`
}

func (Stock) TableName() string {
	return "stocks"
}

func (s *Stock) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

type StockResponse struct {
	ID         uuid.UUID `json:"id"`
	Ticker     string    `json:"ticker"`
	Company    string    `json:"company"`
	Brokerage  string    `json:"brokerage"`
	Action     string    `json:"action"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom float64   `json:"target_from"`
	TargetTo   float64   `json:"target_to"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (s *Stock) ToResponse() StockResponse {
	return StockResponse{
		ID:         s.ID,
		Ticker:     s.Ticker,
		Company:    s.Company,
		Brokerage:  s.Brokerage,
		Action:     s.Action,
		RatingFrom: s.RatingFrom,
		RatingTo:   s.RatingTo,
		TargetFrom: s.TargetFrom,
		TargetTo:   s.TargetTo,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}
