<script setup lang="ts">
import { computed, type FunctionalComponent, type HTMLAttributes, type VNodeProps } from 'vue';
import { useStocksRecommendationsStore } from '@/stores/stocksRecommendations';
import { TrophyIcon, ArrowUpIcon, ArrowDownIcon, ArrowPathIcon } from '@heroicons/vue/24/solid';
import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
import { ref } from 'vue';
import type { StockRecoomendation } from '@/types/stocks';

const stocksRecommendationsStore = useStocksRecommendationsStore();
const isDetailsOpen = ref(false);
const selectedStock = ref<StockRecoomendation | null>(null);

const topThreeStocks = computed(() => {
  return stocksRecommendationsStore.stocksRecommendations.slice(0, 3);
});

const positionColors: { [key: number]: string } = {
  0: 'bg-yellow-500', 
  1: 'bg-gray-400', 
  2: 'bg-amber-700',
};

const actionIcons: { [key: string]: FunctionalComponent<HTMLAttributes & VNodeProps, {}, any, {}> } = {
  'target raised by': ArrowUpIcon,
  'upgraded by': ArrowUpIcon,
  'target lowered by': ArrowDownIcon,
  'downgraded by': ArrowDownIcon,
  'reiterated by': ArrowPathIcon,
  'initiated by': ArrowPathIcon,
};

const actionColors: { [key: string]: string } = {
  'target raised by': 'text-green-600',
  'upgraded by': 'text-green-600',
  'target lowered by': 'text-red-600',
  'downgraded by': 'text-red-600',
  'reiterated by': 'text-blue-600',
  'initiated by': 'text-purple-600',
};

const ratingColors: { [key: string]: string } = {
  'Buy': 'bg-green-100 text-green-800',
  'Outperform': 'bg-green-100 text-green-800',
  'Overweight': 'bg-green-100 text-green-800',
  'Equal Weight': 'bg-blue-100 text-blue-800',
  'Sector Perform': 'bg-blue-100 text-blue-800',
  'Hold': 'bg-yellow-100 text-yellow-800',
  'Neutral': 'bg-yellow-100 text-yellow-800',
  'Underperform': 'bg-red-100 text-red-800',
  'Sell': 'bg-red-100 text-red-800',
};

const showDetails = (stock: StockRecoomendation) => {
  selectedStock.value = stock;
  isDetailsOpen.value = true;
};

const getChangePercentClass = (changePercent: number) => {
  return changePercent >= 0 ? 'text-green-600' : 'text-red-600';
};
</script>

<template>
  <div class="py-6">
    <div class="flex justify-center">
      <h2 class="text-2xl font-bold text-gray-900 flex items-center mb-6">
        <TrophyIcon class="h-8 w-8 text-yellow-500 mr-2" />
        Top Stock Recommendations
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
            style="height: 180px;" @click="showDetails(topThreeStocks[1])">
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
            style="height: 220px;" @click="showDetails(topThreeStocks[0])">
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
            style="height: 140px;" @click="showDetails(topThreeStocks[2])">
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
              <div class="px-4 py-4 sm:px-6 cursor-pointer" @click="showDetails(stock)">
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

    <TransitionRoot appear :show="isDetailsOpen" as="template">
      <Dialog as="div" @close="isDetailsOpen = false" class="relative z-10">
        <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0" enter-to="opacity-100"
          leave="duration-200 ease-in" leave-from="opacity-100" leave-to="opacity-0">
          <div class="fixed inset-0 bg-black bg-opacity-25" />
        </TransitionChild>

        <div class="fixed inset-0 overflow-y-auto">
          <div class="flex min-h-full items-center justify-center p-4">
            <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0 scale-95"
              enter-to="opacity-100 scale-100" leave="duration-200 ease-in" leave-from="opacity-100 scale-100"
              leave-to="opacity-0 scale-95">
              <DialogPanel v-if="selectedStock"
                class="w-full max-w-md transform overflow-hidden rounded-lg bg-white p-6 text-left align-middle shadow-xl transition-all">
                <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                  {{ selectedStock.ticker }} - {{ selectedStock.company }}
                </DialogTitle>

                <div class="mt-4 space-y-4">
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500">Rating</span>
                    <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                      :class="ratingColors[selectedStock.rating]">
                      {{ selectedStock.rating }}
                    </span>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500">Target Price</span>
                    <span class="font-medium">${{ selectedStock.target_price }}</span>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500">Action</span>
                    <div class="flex items-center">
                      <component :is="actionIcons[selectedStock.action]" class="h-4 w-4 mr-1"
                        :class="actionColors[selectedStock.action]" />
                      <span class="text-sm font-medium">{{ selectedStock.action }}</span>
                    </div>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500">Change Percent</span>
                    <span class="font-medium" :class="getChangePercentClass(selectedStock.change_percent)">
                      {{ selectedStock.change_percent >= 0 ? '+' : '' }}{{ selectedStock.change_percent.toFixed(2) }}%
                    </span>
                  </div>

                  <div class="flex items-center justify-between">
                    <span class="text-sm text-gray-500">Score</span>
                    <span class="font-medium">{{ selectedStock.score.toFixed(1) }}</span>
                  </div>
                </div>

                <div class="mt-6">
                  <button type="button"
                    class="w-full inline-flex justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
                    @click="isDetailsOpen = false">
                    Close
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>