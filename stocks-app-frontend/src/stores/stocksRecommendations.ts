import StocksService from "@/services/StocksService";
import type { StockRecoomendation } from "@/types/stocks";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useStocksRecommendationsStore = defineStore('stocksRecommendations', () => {
  const stocksRecommendations = ref<StockRecoomendation[]>([]);
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  async function fetchStockRecommendations() {
    isLoading.value = true;
    error.value = null;

    try {
      const response = await StocksService.getStockRecommendations();

      if (response && response.success) {
        stocksRecommendations.value = response.data;
      } else {
        error.value = 'Failed to fetch stock recommendations data';
      }
    } catch (err) {
      error.value = 'An error occurred while fetching stock recommendations';
      console.error(err);
    } finally {
      isLoading.value = false;
    }
  }

  return {
    stocksRecommendations,
    isLoading,
    error,

    fetchStockRecommendations
  }
});