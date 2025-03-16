export interface Stock {
  id: string;
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  rating_from: string;
  rating_to: string;
  target_from: number;
  target_to: number;
  created_at: string;
  updated_at: string;
}

export interface StockRecoomendation {
  ticker: Stock['ticker'];
  company: Stock['company'];
  score: number;
  rating: string;
  target_price: number;
  action: Stock['action'];
  change_percent: number;
}

export interface StockRecommendationResponse {
  success: boolean;
  message: string;
  data: StockRecoomendation[];
  meta: {
    current_page: number;
    page_size: number;
    total_items: number;
    total_pages: number;
  };
}

export interface StockResponse {
  success: boolean;
  message: string;
  data: Stock[];
  meta: {
    current_page: number;
    page_size: number;
    total_items: number;
    total_pages: number;
  };
}