<script setup lang="ts">
  import { computed, onMounted } from 'vue';
  import { useStocksRecommendationsStore } from '@/stores/stocksRecommendations';
  import { TrophyIcon } from '@heroicons/vue/24/solid';
  import { actionColors, actionIcons, ratingColors, getChangePercentClass } from '@/utils/index';
  import ModalStockRecommendations from '@/components/ModalStockRecommendations.vue';

  const stocksRecommendationsStore = useStocksRecommendationsStore();

  onMounted(() => {
    stocksRecommendationsStore.fetchStockRecommendations();
  });

  const topThreeStocks = computed(() => {
    return stocksRecommendationsStore.stocksRecommendations.slice(0, 3);
  });
</script>

<template>
  <div class="py-6">
    <div class="flex justify-center">
      <h2 class="text-2xl font-bold text-gray-900 flex items-center mb-6">
        <TrophyIcon class="h-8 w-8 text-yellow-500 mr-2" />
        {{ $t('topStocksTitle') }}
      </h2>
    </div>

    <div v-if="stocksRecommendationsStore.isLoading" class="flex justify-center py-12">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
    </div>

    <div v-else-if="stocksRecommendationsStore.error" class="text-center py-8">
      <div class="text-red-600 text-lg">
        {{ stocksRecommendationsStore.error }}
      </div>
      <button @click="stocksRecommendationsStore.fetchStockRecommendations()"
        class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
        Try Again
      </button>
    </div>

    <div v-else-if="topThreeStocks.length > 0" class="flex flex-col items-center">
      <div class="flex items-end justify-center space-x-4 mb-6 w-full max-w-3xl">
        <div v-if="topThreeStocks.length > 1" class="flex flex-col items-center w-1/4">
          <div class="text-center mb-2">
            <span
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-400 text-white">
              2<sup>nd</sup> Place
            </span>
          </div>
          <div
            class="w-full rounded-t-lg bg-gray-400 px-4 pt-4 pb-2 shadow-lg hover:shadow-xl transition-shadow cursor-pointer"
            style="height: 180px;" @click="stocksRecommendationsStore.handleStockClick(topThreeStocks[1])">
            <div class="flex flex-col items-center h-full">
              <span class="text-white font-bold text-xl">{{ topThreeStocks[1].ticker }}</span>
              <span class="text-white text-sm truncate w-full text-center">{{ topThreeStocks[1].company }}</span>
              <div class="mt-auto">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="ratingColors[topThreeStocks[1].rating]">
                  {{ topThreeStocks[1].rating }}
                </span>
              </div>
              <div class="flex items-center justify-center mt-1">
                <component :is="actionIcons[topThreeStocks[1].action]" class="h-4 w-4 mr-1"
                  :class="actionColors[topThreeStocks[1].action]" />
                <span :class="getChangePercentClass(topThreeStocks[1].change_percent)">
                  {{ topThreeStocks[1].change_percent >= 0 ? '+' : '' }}{{ topThreeStocks[1].change_percent.toFixed(2)
                  }}%
                </span>
              </div>
            </div>
          </div>
          <div class="bg-gray-200 text-center py-2 w-full text-gray-800 font-bold rounded-b-lg">
            {{ topThreeStocks[1].score.toFixed(1) }}
          </div>
        </div>

        <!-- First Place (Taller) -->
        <div class="flex flex-col items-center w-1/3">
          <div class="text-center mb-2">
            <span
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-yellow-500 text-white">
              1<sup>st</sup> Place
            </span>
          </div>
          <div
            class="w-full rounded-t-lg bg-yellow-500 px-4 pt-4 pb-2 shadow-lg hover:shadow-xl transition-shadow cursor-pointer"
            style="height: 220px;" @click="stocksRecommendationsStore.handleStockClick(topThreeStocks[0])">
            <div class="flex flex-col items-center h-full">
              <TrophyIcon class="h-8 w-8 text-white mb-1" />
              <span class="text-white font-bold text-2xl">{{ topThreeStocks[0].ticker }}</span>
              <span class="text-white text-sm truncate w-full text-center">{{ topThreeStocks[0].company }}</span>
              <div class="mt-auto">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="ratingColors[topThreeStocks[0].rating]">
                  {{ topThreeStocks[0].rating }}
                </span>
              </div>
              <div class="flex items-center justify-center mt-1">
                <component :is="actionIcons[topThreeStocks[0].action]" class="h-4 w-4 mr-1"
                  :class="actionColors[topThreeStocks[0].action]" />
                <span :class="getChangePercentClass(topThreeStocks[0].change_percent)">
                  {{ topThreeStocks[0].change_percent >= 0 ? '+' : '' }}{{ topThreeStocks[0].change_percent.toFixed(2)
                  }}%
                </span>
              </div>
            </div>
          </div>
          <div class="bg-gray-200 text-center py-2 w-full text-gray-800 font-bold rounded-b-lg">
            {{ topThreeStocks[0].score.toFixed(1) }}
          </div>
        </div>

        <!-- Third Place -->
        <div v-if="topThreeStocks.length > 2" class="flex flex-col items-center w-1/4">
          <div class="text-center mb-2">
            <span
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-amber-700 text-white">
              3<sup>rd</sup> Place
            </span>
          </div>
          <div
            class="w-full rounded-t-lg bg-amber-700 px-4 pt-4 pb-2 shadow-lg hover:shadow-xl transition-shadow cursor-pointer"
            style="height: 140px;" @click="stocksRecommendationsStore.handleStockClick(topThreeStocks[2])">
            <div class="flex flex-col items-center h-full">
              <span class="text-white font-bold text-xl">{{ topThreeStocks[2].ticker }}</span>
              <span class="text-white text-sm truncate w-full text-center">{{ topThreeStocks[2].company }}</span>
              <div class="mt-auto">
                <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                  :class="ratingColors[topThreeStocks[2].rating]">
                  {{ topThreeStocks[2].rating }}
                </span>
              </div>
              <div class="flex items-center justify-center mt-1">
                <component :is="actionIcons[topThreeStocks[2].action]" class="h-4 w-4 mr-1"
                  :class="actionColors[topThreeStocks[2].action]" />
                <span :class="getChangePercentClass(topThreeStocks[2].change_percent)">
                  {{ topThreeStocks[2].change_percent >= 0 ? '+' : '' }}{{ topThreeStocks[2].change_percent.toFixed(2)
                  }}%
                </span>
              </div>
            </div>
          </div>
          <div class="bg-gray-200 text-center py-2 w-full text-gray-800 font-bold rounded-b-lg">
            {{ topThreeStocks[2].score.toFixed(1) }}
          </div>
        </div>
      </div>

      <div class="w-full max-w-3xl mt-8" v-if="stocksRecommendationsStore.stocksRecommendations.length > 3">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Other Recommendations</h3>
        <div class="bg-white shadow overflow-hidden sm:rounded-md">
          <ul class="divide-y divide-gray-200">
            <li v-for="(stock, index) in stocksRecommendationsStore.stocksRecommendations.slice(3)" :key="stock.ticker"
              class="hover:bg-gray-50">
              <div class="px-4 py-4 sm:px-6 cursor-pointer" @click="stocksRecommendationsStore.handleStockClick(stock)">
                <div class="flex items-center justify-between">
                  <div class="flex items-center">
                    <span
                      class="inline-flex items-center justify-center h-8 w-8 rounded-full bg-gray-200 text-gray-600 text-sm font-medium mr-3">
                      {{ index + 4 }}
                    </span>
                    <div>
                      <p class="text-sm font-medium text-blue-600 truncate">
                        {{ stock.ticker }} - {{ stock.company }}
                      </p>
                      <div class="flex items-center mt-1">
                        <span class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium mr-2"
                          :class="ratingColors[stock.rating]">
                          {{ stock.rating }}
                        </span>
                        <component :is="actionIcons[stock.action]" class="h-4 w-4 mr-1"
                          :class="actionColors[stock.action]" />
                        <span :class="getChangePercentClass(stock.change_percent)">
                          {{ stock.change_percent >= 0 ? '+' : '' }}{{ stock.change_percent.toFixed(2) }}%
                        </span>
                      </div>
                    </div>
                  </div>
                  <div class="ml-2 flex-shrink-0">
                    <p class="text-sm text-gray-600">Score: <span class="font-semibold">{{ stock.score.toFixed(1)
                        }}</span></p>
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>

    <div v-else class="text-center py-8">
      <p class="text-gray-500">No recommendations available.</p>
      <button @click="stocksRecommendationsStore.fetchStockRecommendations()"
        class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500">
        Refresh
      </button>
    </div>

    <ModalStockRecommendations />

  </div>
</template>