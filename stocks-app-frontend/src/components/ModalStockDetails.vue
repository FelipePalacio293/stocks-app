<script setup lang="ts">
    import { Dialog, DialogPanel, DialogTitle, TransitionChild, TransitionRoot } from '@headlessui/vue';
    import { useStocksStore } from '@/stores/stocks';
    import { actionColors, actionIcons, ratingColors, getChangePercentClass } from '@/utils/index';
    import { formatDate } from '@/utils/index';

    const stocksStore = useStocksStore();
</script>

<template>
  <TransitionRoot appear :show="stocksStore.showStockDetails" as="template">
    <Dialog as="div" @close="stocksStore.closeStockDetails" class="relative z-10">
      <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0" enter-to="opacity-100"
        leave="duration-200 ease-in" leave-from="opacity-100" leave-to="opacity-0">
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4">
          <TransitionChild as="template" enter="duration-300 ease-out" enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100" leave="duration-200 ease-in" leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95">
            <DialogPanel v-if="stocksStore.selectedStock"
              class="w-full max-w-md transform overflow-hidden rounded-lg bg-white p-6 text-left align-middle shadow-xl transition-all">
              <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                {{ stocksStore.selectedStock.ticker }} - {{ stocksStore.selectedStock.company }}
              </DialogTitle>

              <div class="mt-4 space-y-4">
                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Brokerage</span>
                  <span class="font-medium">{{ stocksStore.selectedStock.brokerage }}</span>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Action</span>
                  <div class="flex items-center">
                    <component :is="actionIcons[stocksStore.selectedStock.action]" class="h-4 w-4 mr-1"
                      :class="actionColors[stocksStore.selectedStock.action]" />
                    <span class="text-sm font-medium">{{ stocksStore.selectedStock.action }}</span>
                  </div>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Rating</span>
                  <div class="flex items-center">
                    <span class="text-sm text-gray-500 mr-2">
                      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="ratingColors[stocksStore.selectedStock.rating_from]">
                        {{ stocksStore.selectedStock.rating_from }}
                      </span>
                    </span>
                    <span class="text-sm">→</span>
                    <span class="text-sm ml-2">
                      <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
                        :class="ratingColors[stocksStore.selectedStock.rating_to]">
                        {{ stocksStore.selectedStock.rating_to }}
                      </span>
                    </span>
                  </div>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Target Price</span>
                  <div class="flex items-center">
                    <span class="text-sm mr-2">${{ stocksStore.selectedStock.target_from }}</span>
                    <span class="text-sm">→</span>
                    <span class="text-sm ml-2">${{ stocksStore.selectedStock.target_to }}</span>
                    <span 
                      class="ml-2 text-xs"
                      :class="getChangePercentClass((stocksStore.selectedStock.target_to - stocksStore.selectedStock.target_from) / stocksStore.selectedStock.target_from * 100)"
                    >
                      ({{ (stocksStore.selectedStock.target_to - stocksStore.selectedStock.target_from) >= 0 ? '+' : '' }}{{ ((stocksStore.selectedStock.target_to - stocksStore.selectedStock.target_from) / stocksStore.selectedStock.target_from * 100).toFixed(2) }}%)
                    </span>
                  </div>
                </div>

                <div class="flex items-center justify-between">
                  <span class="text-sm text-gray-500">Updated At</span>
                  <span class="font-medium">{{ formatDate(stocksStore.selectedStock.updated_at) }}</span>
                </div>
              </div>

              <div class="mt-6">
                <button type="button"
                  class="w-full inline-flex justify-center rounded-md border border-transparent bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
                  @click="stocksStore.closeStockDetails">
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