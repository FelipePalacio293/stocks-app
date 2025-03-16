<script setup lang="ts">
  import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
  import { useStocksRecommendationsStore } from '@/stores/stocksRecommendations';
  import { actionColors, actionIcons, ratingColors, getChangePercentClass } from '@/utils/index';

  const stocksRecommendationsStore = useStocksRecommendationsStore();
</script>

<template>
  <TransitionRoot appear :show="stocksRecommendationsStore.showModal" as="template">
    <Dialog as="div" @close="stocksRecommendationsStore.closeModal" class="relative z-10">
      <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0" enter-to="opacity-100"
        leave="duration-200 ease-in" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100" leave="duration-200 ease-in" leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95">
            <DialogPanel v-if="stocksRecommendationsStore.selectedStock"
              class="w-full max-w-md transform overflow-hidden rounded-lg bg-white p-6 text-left align-middle shadow-xl transition-all">
              <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                {{ stocksRecommendationsStore.selectedStock.ticker }} - {{ stocksRecommendationsStore.selectedStock.company }}
              </DialogTitle>

              <div class="mt-4 space-y-4">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Rating</span>
                  <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                    :class="ratingColors[stocksRecommendationsStore.selectedStock.rating]">
                    {{ stocksRecommendationsStore.selectedStock.rating }}
                  </span>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Target Price</span>
                  <span class="font-medium">${{ stocksRecommendationsStore.selectedStock.target_price }}</span>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Action</span>
                  <div class="flex items-center">
                    <component :is="actionIcons[stocksRecommendationsStore.selectedStock.action]" class="h-4 w-4 mr-1"
                      :class="actionColors[stocksRecommendationsStore.selectedStock.action]" />
                    <span class="text-sm font-medium">{{ stocksRecommendationsStore.selectedStock.action }}</span>
                  </div>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Change Percent</span>
                  <span class="font-medium" :class="getChangePercentClass(stocksRecommendationsStore.selectedStock.change_percent)">
                    {{ stocksRecommendationsStore.selectedStock.change_percent >= 0 ? '+' : '' }}{{ stocksRecommendationsStore.selectedStock.change_percent.toFixed(2) }}%
                  </span>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Score</span>
                  <span class="font-medium">{{ stocksRecommendationsStore.selectedStock.score.toFixed(1) }}</span>
                </div>
              </div>

              <div class="mt-6">
                <button type="button"
                  class="w-full inline-flex justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
                  @click="stocksRecommendationsStore.closeModal">
                  Close
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
