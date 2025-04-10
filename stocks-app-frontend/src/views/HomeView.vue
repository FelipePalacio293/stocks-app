<script setup lang="ts">
  import { useStocksStore } from '@/stores/stocks';
  import { onMounted } from 'vue';
  import Stock from '@/components/Stock.vue';
  import Spinner from '@/components/Spinner.vue';
  import ModalStockDetails from '@/components/ModalStockDetails.vue';
  import StockSearch from '@/components/StockSearch.vue';

  const stocksStore = useStocksStore();

  onMounted(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const tickerParam = urlParams.get('ticker');
    
    if (tickerParam) {
      stocksStore.setSearchTerm(tickerParam);
    } else {
      stocksStore.fetchStocks();
    }
  });
</script>

<template>
  <div>
    <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
      <h1 class="text-3xl font-bold tracking-tight text-gray-900">{{ $t('analystTitle') }}</h1>

      <StockSearch />

      <div v-if="stocksStore.isLoading" class="mt-5">
        <Spinner />
      </div>

      <div v-else-if="stocksStore.error" class="mt-5 text-red-500">
        {{ stocksStore.error }}
      </div>

      <div v-else-if="stocksStore.hasStocks" class="mt-8">
        <ul role="list" class="grid grid-cols-1 gap-x-6 gap-y-8 lg:grid-cols-3 xl:gap-x-8">
          <li 
            v-for="stock in stocksStore.stocks" 
            :key="stock.id"
            class="overflow-hidden rounded-xl border border-gray-200"
          >
            <Stock
              :stock="stock"
              :key="stock.id"
            />
          </li>
        </ul>

        <div class="mt-8 flex items-center justify-between">
          <div class="text-sm text-gray-700">
            {{ $t('pagination.page') }} {{ stocksStore.currentPage }} {{ $t('pagination.of') }} {{ stocksStore.totalPages }}
            <span class="text-gray-500">({{ stocksStore.totalItems }} stocks)</span>
          </div>

          <div class="flex space-x-2">
            <button @click="stocksStore.fetchStocks(stocksStore.currentPage - 1)"
              :disabled="stocksStore.currentPage === 1"
              class="px-4 py-2 bg-white border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 disabled:opacity-50">
              {{ $t('pagination.previous') }}
            </button>
            <button @click="stocksStore.fetchStocks(stocksStore.currentPage + 1)"
              :disabled="stocksStore.currentPage === stocksStore.totalPages"
              class="px-4 py-2 bg-indigo-600 border border-transparent rounded-md shadow-sm text-sm font-medium text-white hover:bg-indigo-700 disabled:opacity-50">
              {{ $t('pagination.next') }}
            </button>
          </div>
        </div>
      </div>

      <div v-else class="mt-5">
        <p class="text-gray-500">No stocks found.</p>
      </div>
    </div>
    <ModalStockDetails />
  </div>
</template>
