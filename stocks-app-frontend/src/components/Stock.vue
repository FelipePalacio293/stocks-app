<script setup lang="ts">
  import { Menu, MenuButton, MenuItem, MenuItems } from '@headlessui/vue';
  import { EllipsisHorizontalIcon } from '@heroicons/vue/20/solid'
  import { formatDate } from '@/utils/index';
  import type { Stock } from '../types/stocks';

  defineProps<{
    stock: Stock
  }>();

  const getActionColor = (action: string) => {
    switch (action) {
      case 'target raised by':
        return 'text-green-700 bg-green-50 ring-green-600/20';
      case 'target lowered by':
        return 'text-red-700 bg-red-50 ring-red-600/10';
      case 'reiterated by':
        return 'text-blue-700 bg-blue-50 ring-blue-600/20';
      case 'downgraded by':
        return 'text-orange-700 bg-orange-50 ring-orange-600/20';
      case 'upgraded by':
        return 'text-emerald-700 bg-emerald-50 ring-emerald-600/20';
      default:
        return 'text-gray-600 bg-gray-50 ring-grascript10';
    }
  };

  const getTickerBgColor = (ticker: string) => {
    const hash = ticker.split('').reduce((acc, char) => {
      return char.charCodeAt(0) + ((acc << 5) - acc);
    }, 0);

    const h = Math.abs(hash % 360);
    return `hsl(${h}, 70%, 40%)`;
  };
</script>

<template>
  <div class="flex items-center gap-x-4 border-b border-gray-900/5 bg-gray-50 p-6">
    <div class="flex-none size-12 rounded-lg flex items-center justify-center text-white font-bold"
      :style="{ backgroundColor: getTickerBgColor(stock.ticker) }">
      {{ stock.ticker.substring(0, 2) }}
    </div>

    <div>
      <div class="text-sm font-medium text-gray-900">{{ stock.ticker }}</div>
      <div class="text-xs text-gray-500">{{ stock.company }}</div>
    </div>

    <Menu as="div" class="relative ml-auto">
      <MenuButton class="-m-2.5 block p-2.5 text-gray-400 hover:text-gray-500">
        <span class="sr-only">Open options</span>
        <EllipsisHorizontalIcon class="size-5" aria-hidden="true" />
      </MenuButton>
      <transition enter-active-class="transition ease-out duration-100" enter-from-class="transform opacity-0 scale-95"
        enter-to-class="transform opacity-100 scale-100" leave-active-class="transition ease-in duration-75"
        leave-from-class="transform opacity-100 scale-100" leave-to-class="transform opacity-0 scale-95">
        <MenuItems
          class="absolute right-0 z-10 mt-0.5 w-32 origin-top-right rounded-md bg-white py-2 shadow-lg ring-1 ring-gray-900/5 focus:outline-none">
          <MenuItem v-slot="{ active }">
          <button 
            :class="[active ? 'bg-gray-50' : '', 'block px-3 py-1 text-sm text-gray-900']"
          >
            View Details
          </button>
          </MenuItem>
        </MenuItems>
      </transition>
    </Menu>
  </div>

  <dl class="-my-3 divide-y divide-gray-100 px-6 py-4 text-sm">
    <div class="flex justify-between gap-x-4 py-3">
      <dt class="text-gray-500">Brokerage</dt>
      <dd class="text-gray-700">{{ stock.brokerage }}</dd>
    </div>

    <div class="flex justify-between gap-x-4 py-3">
      <dt class="text-gray-500">Rating</dt>
      <dd class="text-gray-700">
        {{ stock.rating_from }} → {{ stock.rating_to }}
      </dd>
    </div>

    <div class="flex justify-between gap-x-4 py-3">
      <dt class="text-gray-500">Price Target</dt>
      <dd class="flex items-center gap-x-2">
        <div class="font-medium text-gray-900">${{ stock.target_from }} → ${{ stock.target_to }}</div>
        <div :class="[getActionColor(stock.action), 'rounded-md px-2 py-1 text-xs font-medium ring-1 ring-inset']">
          {{ stock.action }}
        </div>
      </dd>
    </div>

    <div class="flex justify-between gap-x-4 py-3">
      <dt class="text-gray-500">Updated</dt>
      <dd class="text-gray-700">
        <time :datetime="stock.updated_at">{{ formatDate(stock.updated_at) }}</time>
      </dd>
    </div>
  </dl>
</template>
