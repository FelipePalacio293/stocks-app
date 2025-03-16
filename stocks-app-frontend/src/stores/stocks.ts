import StocksService from "@/services/StocksService";
import type { Stock } from "@/types/stocks";
import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useStocksStore = defineStore('stocks', () => {
  const stocks = ref<Stock[]>([]);
  const currentPage = ref<number>(1);
  const totalPages = ref<number>(0);
  const totalItems = ref<number>(0);
  const pageSize = ref<number>(10);
  const isLoading = ref<boolean>(false);
  const error = ref<string | null>(null);

  async function fetchStocks(page: number = 1) {
    isLoading.value = true;
    error.value = null;
    
    try {
      const response = await StocksService.getStocks(page);
      
      if (response && response.success) {
        stocks.value = response.data;
        currentPage.value = response.meta.current_page;
        totalPages.value = response.meta.total_pages;
        totalItems.value = response.meta.total_items;
        pageSize.value = response.meta.page_size;
      } else {
        error.value = 'Failed to fetch stocks data';
      }
    } catch (err) {
      error.value = 'An error occurred while fetching stocks';
      console.error(err);
    } finally {
      isLoading.value = false;
    }
  }

  const hasStocks = computed(() => stocks.value.length > 0);

  return {
    stocks,
    currentPage,
    totalPages,
    totalItems,
    pageSize,
    isLoading,
    error,
    
    fetchStocks,
    
    hasStocks
  }
});