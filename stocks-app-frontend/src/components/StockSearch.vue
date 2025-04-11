<script setup lang="ts">
  import { ref, watch } from 'vue';
  import { useStocksStore } from '@/stores/stocks';

  const stocksStore = useStocksStore();
  const searchInput = ref(stocksStore.searchTerm);

  const debouncedSearch = ref(null as any);

  function handleSearch() {
    clearTimeout(debouncedSearch.value);
    debouncedSearch.value = setTimeout(() => {
      stocksStore.setSearchTerm(searchInput.value);
    }, 500);
  }

  watch(() => stocksStore.searchTerm, (newValue) => {
    searchInput.value = newValue;
  });
</script>

<template>
  <div class="my-6">
    <div class="relative rounded-md shadow-sm max-w-3xl mx-auto">
      <div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
        <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
          <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
        </svg>
      </div>
      <input
        type="text"
        v-model="searchInput"
        @input="handleSearch"
        class="block w-full rounded-md border-gray-300 pl-10 focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm h-12 text-base"
        :placeholder="$t('searchPlaceholder') || 'Search by ticker symbol...'"
      />
    </div>
  </div>
</template>
