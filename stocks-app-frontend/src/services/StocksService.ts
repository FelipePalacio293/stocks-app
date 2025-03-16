import type { StockResponse } from '@/types/stocks';
import api from '../lib/axios';

export default {
  async getStocks(page: number = 1): Promise<StockResponse | undefined> {
    try {
      const response = await api.get<StockResponse>(`/api/v1/stocks?page=${page}`);
      return response.data;
    } catch (error) {
      console.error('Error fetching stocks:', error);
      return undefined;
    }
  }
}