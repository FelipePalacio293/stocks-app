<script setup>
  import { Disclosure, DisclosureButton, DisclosurePanel } from '@headlessui/vue'
  import { Bars3Icon, XMarkIcon } from '@heroicons/vue/24/outline'
  import { RouterLink } from 'vue-router'

  const navigation = [
    { name: 'Home', href: '/' },
    { name: 'Stocks recommendations', href: '/stocks' },
  ]
</script>

<template>
    <Disclosure as="nav" class="border-b border-gray-200 bg-gray-100" v-slot="{ open }">
      <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 justify-between">
          <div class="flex">
            <div class="flex shrink-0 items-center">
              <img class="block h-8 w-auto lg:hidden"
                src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=indigo&shade=600"
                alt="Stocks img" 
              />
              <img class="hidden h-8 w-auto lg:block"
                src="https://tailwindcss.com/plus-assets/img/logos/mark.svg?color=indigo&shade=600"
                alt="Stocks img" 
              />
            </div>
            <div class="hidden sm:-my-px sm:ml-6 sm:flex sm:space-x-8">
              <RouterLink 
                v-for="item in navigation" 
                :key="item.name" 
                :to="item.href"
                :class="['border-transparent font-bold text-gray-700 hover:border-gray-300 hover:text-gray-700', 'inline-flex items-center border-b-2 px-1 pt-1 text-sm']"
                active-class="border-indigo-500 font-extrabold text-gray-900"
              >
                {{ item.name }}
              </RouterLink>
            </div>
          </div>
          <div class="-mr-2 flex items-center sm:hidden">
            <DisclosureButton
              class="relative inline-flex items-center justify-center rounded-md bg-white p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-hidden">
              <span class="absolute -inset-0.5" />
              <span class="sr-only">Open main menu</span>
              <Bars3Icon v-if="!open" class="block size-6" aria-hidden="true" />
              <XMarkIcon v-else class="block size-6" aria-hidden="true" />
            </DisclosureButton>
          </div>
        </div>
      </div>

      <DisclosurePanel class="sm:hidden">
        <div class="space-y-1 pt-2 pb-3">
          <DisclosureButton v-for="item in navigation" :key="item.name" as="a" :href="item.href"
            :class="[item.current ? 'border-indigo-500 bg-indigo-50 text-indigo-700' : 'border-transparent text-gray-600 hover:border-gray-300 hover:bg-gray-50 hover:text-gray-800', 'block border-l-4 py-2 pr-4 pl-3 text-base font-medium']"
            :aria-current="item.current ? 'page' : undefined">{{ item.name }}</DisclosureButton>
        </div>
      </DisclosurePanel>
    </Disclosure>
</template>
