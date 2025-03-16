import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
    },
    {
      path: '/stocks',
      name: 'Stocks',
      component: () => import('@/views/StockRecommendationsView.vue'),
    }
  ],
})

export default router
