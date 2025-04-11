import type { StockRecommendationResponse, StockResponse } from '@/types/stocks';
import api from '../lib/axios';

export default {
  async getStocks(page: number = 1, ticker: string): Promise<StockResponse | undefined> {
    try {
      const params = new URLSearchParams();
      if (ticker) {
        params.append('ticker', ticker);
      }

      const response = await api.get<StockResponse>(`/api/v1/stocks?page=${page}&page_size=12&${params.toString()}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching stocks:', error);
      return undefined;
    }
  },
  async getStockRecommendations(): Promise<StockRecommendationResponse | undefined> {
    try {
      const response = await api.get<StockRecommendationResponse>(`/api/v1/stocks/recommendations`);
      return response.data;
    } catch (error) {
      console.error('Error fetching stock recommendations:', error);
      return undefined;
    }
  },
}